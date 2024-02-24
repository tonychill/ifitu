package redis

import (
	"context"
	"fmt"
	"strings"

	r9 "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/tonychill/ifitu/apis/pb/go/global"
	"github.com/tonychill/ifitu/lib/utils"
)

func (c *clientImpl) HasIndex(ctx context.Context, index string) (bool, error) {

	resp, err := c.conn.Do(ctx, "FT._LIST").Result()
	if err != nil {
		return false, err
	}

	for _, idx := range resp.([]interface{}) {
		if strings.Contains(idx.(string), index) {
			return true, nil
		}
	}

	return false, nil

}

// TODO: should deprecate and make this internal to the package as external users should use
// methods that are specific to their use case such as Search and Index
func (c *clientImpl) Do(ctx context.Context, args ...interface{}) *r9.Cmd {
	return c.conn.Do(ctx, args...)
}

func (c *clientImpl) JsonAppend(ctx context.Context, ap Append) (err error) {
	args := Args{"JSON.ARRAPPEND"}
	if ap.Index != "" {
		args = append(args, ap.Index+":"+ap.Key)
	} else {
		args = append(args, c.name+":"+ap.Key)
	}
	args = append(args, "$."+ap.Field)
	for _, v := range ap.Values {
		args = append(args, fmt.Sprintf(`'"%s"'`, v))
	}
	resp, err := c.conn.Do(ctx, args...).Result()
	if err != nil {
		log.Error().Err(err).Msg("error executing search")
		return err
	}

	log.Debug().Msgf("*** TESTING: resp from appending records: %v", resp)

	return
}

// When saving a json document be sure that all the fields match the expected type
// based on the schema. If a value's type is not the same as defined in the schema
// then that record will not be indexed by the search module.
func (c *clientImpl) CreateJSON(ctx context.Context, req CreateJsonRequest) (id string, err error) {
	if req.Prefix == "" {
		return "", fmt.Errorf("the prefix in the request was empty")
	}

	if req.Index == "" {
		return "", fmt.Errorf("the index in the request was empty")

	}

	if req.Id == "" {
		req.Id, err = utils.NewULID(req.Prefix)
		if err != nil {
			return "", err
		}
	}

	b, err := utils.EncodeV2(req.Object, utils.EncodingType_JSON)
	if err != nil {
		log.Error().Err(err).Msgf("error encoding data: %T", req.Object)
		return "", err
	}

	if _, err := c.conn.Do(ctx, "JSON.SET", SetRedisKey(req.Index, req.Id),
		"$", string(b), req.Args).Result(); err != nil {
		log.Error().Err(err).
			Msg("error saving data to redis when creating a resource.")
		return req.Id, err
	}
	return req.Id, nil

	// args := Args{"JSON.SET"}
	// if req.Index != "" {
	// 	args = append(args, req.Index+":"+req.Id)
	// } else {
	// 	args = append(args, c.name+":"+req.Id)
	// }
	// args = append(args, "$", req.Object)
	// _, err = c.conn.Do(ctx, args...).Result()
	// if err != nil {
	// 	log.Error().Err(err).Msg("error executing search")
	// 	return err
	// }
	// return

}

// TODO: move the querying logic to the lib
func (c *clientImpl) GetJSON(ctx context.Context, query *global.Query, msg interface{}) ([]interface{}, error) {

	if err := utils.ValidateQuery(query); err != nil {
		log.Error().Err(err).Msg("the provided query was not valid")
		return nil, fmt.Errorf("the provided query was not valid")
	}

	var (
	// err error
	// val interface{}
	)

	// q := NewQuery(query).
	// 	Limit(int(query.Page), int(query.PageSize))
	var all bool
	for _, term := range query.Terms {
		if term.Key == "all" {
			all = true
		}

	}

	if all {

		// vals := []*pb.Resource{}
		// val, err = c.conn.Do(ctx, "JSON.GEt", r.redisConfig.Index+":*").Result()
		// if err != nil {
		// 	return nil, err
		// }
		// list, ok := val.([]interface{})
		// if ok {
		// 	// if len(list) == 0 {
		// 	// 	if err := r.reloadDomain(ctx, nil); err != nil {
		// 	// 			Msg("error reloading domain when getting all operations")
		// 	// 		return nil, err
		// 	// 	}
		// 	// 	// Try to get all the keys again after reloading the domain
		// 	// 	val, err = r.redisClient.Do(ctx, "KEYS", r.redisConfig.Index+":*").Result()
		// 	// 	if err != nil {
		// 	// 		return nil, err
		// 	// 	}
		// 	// 	list, ok = val.([]interface{})
		// 	// 	if !ok {
		// 	// 			Msg("The response from redis did not return a list event after reloading the domain.")
		// 	// 		return nil, status.Error(codes.Internal, "The response from redis did not return a list.")
		// 	// 	}
		// 	// }
		// 	for _, key := range list {
		// 		k, _ok := key.(string)
		// 		internalId := strings.ReplaceAll(k, r.redisConfig.Index+":", "")
		// 		if _ok {
		// 			_operations, err := r.getResourcesJSON(ctx, []string{internalId}, false)
		// 			if err != nil {
		// 				log.Error().Msg("error getting operations when tring to get them all.")
		// 				// 	Msg("error getting operations when tring to get them all.")
		// 			}
		// 			operations = append(operations, _operations...)
		// 		} else {
		// 			// 	Msgf("didn't find a json string for id: %s when getting operations json",
		// 			// 		internalId)
		// 			continue
		// 		}
		// 	}
		// } else {
		// 	// 	Msgf("The response from redis did not return a list | resp: type = %T, val = %v",
		// 	// 		val, val)
		// 	return operations, nil //status.Error(codes.Internal, "The response from redis did not return a list.")
		// }
		// return operations, nil
	} else {
		// operations := make([]*pb.Resource, 0)
		// for _, id := range ids {
		// 	if id == "" {
		// 		continue
		// 	}
		// 	// TODO: collect the error and return them when all the calls are done.
		// 	res, err := r.redisClient.Do(ctx, "JSON.GET", redis.SetRedisKey(r.redisConfig.Index, id), "$").Result()
		// 	if err != nil {
		// 		log.Error().Msgf("error making call to redis when getting operations for search key %s", id)
		// 		// 	Msgf("error making call to redis when getting operations for search key %s",
		// 		// 		r.redisConfig.Index+":"+id)
		// 		continue
		// 	}
		// 	operationsJsonStr, ok := res.(string)
		// 	if !ok {
		// 		// 	Msg("The response from redis did not return a json string.")
		// 		continue
		// 	}

		// 	operationsJsonStr = strings.Replace(operationsJsonStr, "[", "", 1)
		// 	operationsJsonStr = strings.Replace(operationsJsonStr, "]", "", 1)
		// 	resource := &pb.Resource{}
		// 	if err := utils.Decode(operationsJsonStr, resource, utils.EncodingType_JSON); err != nil {
		// 		// 	Msgf("error decoding json from redis operations for id %s", id)
		// 		continue
		// 	}

		// 	operations = append(operations, resource)
		// }
		// return operations, nil
	}
	return nil, nil
}

// deprecated: If a json object is saved and the
func (c *clientImpl) Search(ctx context.Context, query *global.Query) (docs []Document, total int64, err error) {
	if err := utils.ValidateQuery(query); err != nil {
		log.Error().Err(err).Msg("the provided query was not valid")
		return nil, 0, fmt.Errorf("the provided query was not valid")
	}
	q := NewQuery(query).
		Limit(int(query.Page), int(query.PageSize))
	args := Args{"FT.SEARCH", c.name}
	args = append(args, q.serialize()...)
	resp, err := c.conn.Do(ctx, args...).Result()
	if err != nil {
		log.Error().Err(err).Msg("error executing search")
		return nil, 0, err
	}

	vals, reply, total, err := values(resp, err)
	if err != nil {
		log.Error().Err(err).Msg("error parsing valuse from response")
		return
	}

	if reply.Results != nil {
		docs = reply.ToDocs()
	} else if len(vals) > 1 {
		docs = make([]Document, 0, len(vals)-1)
	}

	skip := 1
	scoreIdx := -1
	fieldsIdx := -1
	payloadIdx := -1
	if q.Flags&QueryWithScores != 0 {
		scoreIdx = 1
		skip++
	}
	if q.Flags&QueryWithPayloads != 0 {
		payloadIdx = skip
		skip++
	}

	if q.Flags&QueryNoContent == 0 {
		fieldsIdx = skip
		skip++
	}

	if len(docs) == 0 && len(vals) > skip {
		for idIdx := 1; idIdx < len(vals); idIdx += skip {
			if d, e := loadDocument(vals, idIdx, scoreIdx, payloadIdx, fieldsIdx); e == nil {
				docs = append(docs, d)
			} else {
				log.Print("Error parsing doc: ", e)
			}
		}
	}

	return
}

func (c *clientImpl) SearchV2(ctx context.Context, query *global.Query, msg interface{}) (recs []interface{}, total int64, err error) {
	if err := utils.ValidateQuery(query); err != nil {
		log.Error().Err(err).Msg("the provided query was not valid")
		return nil, 0, fmt.Errorf("the provided query was not valid")
	}
	q := NewQuery(query).
		Limit(int(query.Page), int(query.PageSize))
	args := Args{"FT.SEARCH", c.name}
	args = append(args, q.serialize()...)
	resp, err := c.conn.Do(ctx, args...).Result()
	if err != nil {
		log.Error().Err(err).Msg("error executing search")
		return nil, 0, err
	}

	vals, reply, total, err := values(resp, err)
	if err != nil {
		log.Error().Err(err).Msg("error parsing valuse from response")
		return
	}
	var docs []Document
	if reply.Results != nil {
		docs = reply.ToDocs()
	} else if len(vals) > 1 {
		docs = make([]Document, 0, len(vals)-1)
	}

	skip := 1
	scoreIdx := -1
	fieldsIdx := -1
	payloadIdx := -1
	if q.Flags&QueryWithScores != 0 {
		scoreIdx = 1
		skip++
	}
	if q.Flags&QueryWithPayloads != 0 {
		payloadIdx = skip
		skip++
	}

	if q.Flags&QueryNoContent == 0 {
		fieldsIdx = skip
		skip++
	}

	if len(docs) == 0 && len(vals) > skip {
		for idIdx := 1; idIdx < len(vals); idIdx += skip {
			if d, e := loadDocument(vals, idIdx, scoreIdx, payloadIdx, fieldsIdx); e == nil {
				docs = append(docs, d)
			} else {
				log.Print("Error parsing doc: ", e)
			}
		}
	}

	for _, doc := range docs {
		if err := doc.Decode(msg); err != nil {
			log.Error().Err(err).Msg("error converting redis document to rates to proto.")
			return nil, 0, err
		}
		recs = append(recs, msg)
	}

	return
}

// IndexOptions indexes multiple documents on the index, with optional Options passed to options
func (c *clientImpl) IndexOptions(ctx context.Context, opts IndexingOptions, docs ...Document) error {

	n := 0
	var merr MultiError

	for ii, doc := range docs {
		args := make(Args, 0, 6+len(doc.Properties))
		args = append(args, "FT.ADD", c.name, doc.Id, doc.Score)
		args = SerializeIndexingOptions(opts, args)

		if doc.Payload != nil {
			args = args.Add("PAYLOAD", doc.Payload)
		}

		args = append(args, "FIELDS")

		for k, f := range doc.Properties {
			args = append(args, k, f)
		}

		if _, err := c.conn.Do(ctx, args...).Result(); err != nil {
			if merr == nil {
				merr = NewMultiError(len(docs))
			}
			merr[ii] = err

			return merr
		}
		n++
	}

	// TODO: validate the following commented out code.
	// if err := conn.Flush(); err != nil {
	// 	return err
	// }

	// for n > 0 {
	// 	if _, err := conn.Receive(); err != nil {
	// 		if merr == nil {
	// 			merr = NewMultiError(len(docs))
	// 		}
	// 		merr[n-1] = err
	// 	}
	// 	n--
	// }

	if merr == nil {
		return nil
	}

	return merr
}
