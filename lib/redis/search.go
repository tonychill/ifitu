package redis

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// NewDocument creates a document with the specific id and score
func NewDocument(arr []interface{}, idIdx int, score float32) Document {
	var id string
	switch arr[idIdx].(type) {
	case []byte:
		id = string(arr[idIdx].([]byte))
	case string:
		id = arr[idIdx].(string)
	}
	return Document{
		Id:         id,
		Score:      score,
		Properties: make(map[string]interface{}),
	}
}

// SetPayload Sets the document payload
func (d *Document) SetPayload(payload []byte) {
	d.Payload = payload
}

// Set sets a property and its value in the document
func (d Document) Set(name string, value interface{}) Document {
	d.Properties[name] = value
	return d
}

// All punctuation marks and whitespaces (besides underscores) separate the document and queries into tokens.
// e.g. any character of `,.<>{}[]"':;!@#$%^&*()-+=~` will break the text into terms.
// So the text `foo-bar.baz...bag` will be tokenized into `[foo, bar, baz, bag]`
// Escaping separators in both queries and documents is done by prepending a backslash to any separator.
// e.g. the text `hello\-world hello-world` will be tokenized as `[hello-world, hello, world]`.
// **NOTE** that in most languages you will need an extra backslash when formatting the document or query,
// to signify an actual backslash, so the actual text in redis-cli for example, will be entered as `hello\\-world`.
// Underscores (`_`) are not used as separators in either document or query.
// So the text `hello_world` will remain as is after tokenization.
func EscapeTextFileString(value string) string {
	for _, char := range field_tokenization {
		value = strings.Replace(value, string(char), ("\\" + string(char)), -1)
	}
	return value
}

// convert the result from a redis query to a proper Document object
func loadDocument(arr []interface{}, idIdx, scoreIdx, payloadIdx, fieldsIdx int) (Document, error) {

	var score float64 = 1
	var err error
	if scoreIdx > 0 {
		if score, err = strconv.ParseFloat(string(arr[idIdx+scoreIdx].([]byte)), 64); err != nil {
			return Document{}, fmt.Errorf("could not parse score: %s", err)
		}
	}
	doc := NewDocument(arr, idIdx, float32(score))

	if payloadIdx > 0 {
		doc.Payload, _ = arr[idIdx+payloadIdx].([]byte)
	}

	if fieldsIdx > 0 {
		lst := arr[idIdx+fieldsIdx].([]interface{})
		doc.loadFields(lst)
	}

	return doc, nil
}

// SetPayload Sets the document payload
func (d *Document) loadFields(lst []interface{}) *Document {
	for i := 0; i < len(lst); i += 2 {
		var prop string
		switch lst[i].(type) {
		case []byte:
			prop = string(lst[i].([]byte))
		default:
			prop = lst[i].(string)
		}

		var val interface{}
		switch v := lst[i+1].(type) {
		case []byte:
			val = string(v)
		default:
			val = v
		}
		*d = d.Set(prop, val)
	}
	return d
}

type redisReply struct {
	TotalResults int64         `json:"total_results"`
	Results      []Result      `json:"results"`
	Format       string        `json:"format"`
	Error        []interface{} `json:"error"`
	Attibutes    []interface{} `json:"attributes"`
	Values       []interface{} `json:"values"`
}

func (r redisReply) String() string {
	return fmt.Sprintf("TotalResults: %d | Results: %v | Format: %s | Error: %v | Attibutes: %v | Values: %v",
		r.TotalResults, r.Results, r.Format, r.Error, r.Attibutes, r.Values)
}

func (r redisReply) ToDocs() []Document {
	docs := make([]Document, 0, len(r.Results))
	for _, result := range r.Results {
		docs = append(docs, Document{
			Id:         result.Key,
			Score:      1,
			Properties: map[string]interface{}{"$": result.Value},
		})
	}
	return docs
}

type Result struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	// Values interface{} `json:"values"`
}

func values(raw interface{}, err error) ([]interface{}, redisReply, int64, error) {
	reply := redisReply{}
	if err != nil {
		return nil, reply, 0, err
	}

	switch rawReplyCast := raw.(type) {
	case map[interface{}]interface{}:
		ttlRes, _ := rawReplyCast["total_results"].(int64)
		format, _ := rawReplyCast["format"].(string)
		err, _ := rawReplyCast["error"].([]interface{})
		att, _ := rawReplyCast["attributes"].([]interface{})
		final := redisReply{
			TotalResults: ttlRes,
			Format:       format,
			Attibutes:    att,
			Error:        err,
		}

		results, ok := rawReplyCast["results"].([]interface{})
		if !ok {
			log.Error().Msg("the response did not have the expected key of 'results'")
			// return nil, 0, redisReply, fmt.Errorf("the response did not have a key of results")
		}
		for _, result := range results {
			tempFinalRes := Result{}
			switch r := result.(type) {
			case map[interface{}]interface{}:
				key, ok := r["id"].(string)
				if !ok {
					log.Error().Msg("the response did not have a key of 'id'")
				}
				tempFinalRes.Key = key
				// TODO: may implement in the future as its a property that is returned by redis
				// vals, ok := r["values"].([]interface{})
				// if !ok {
				// log.Error().Msg("the response did not have a key of 'id'")

				// }

				ea, ok := r["extra_attributes"].(map[interface{}]interface{})
				if !ok {
					log.Error().Msg("the response did not have a key of 'extra_attributes'")
				}
				val, ok := ea["$"].(string)
				if !ok {
					log.Error().Msg("the response did not have a key of '$'")
				}
				tempFinalRes.Value = val

			}
			final.Results = append(final.Results, tempFinalRes)
		}

		return nil, final, final.TotalResults, nil
	case []interface{}:
		return rawReplyCast, reply, 0, nil
	case nil:
		return nil, reply, 0, ErrNil
	case Error:
		return nil, reply, 0, fmt.Errorf("%s", rawReplyCast) // FIXME: should we return an error here?
	}
	return nil, reply, 0, fmt.Errorf("unexpected type for Values, got type %T", reply)
}
