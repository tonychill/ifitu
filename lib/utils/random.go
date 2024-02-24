package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"strings"
	"time"

	// "go.starlark.net/lib/proto"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid"
	"github.com/tonychill/ifitu/apis/pb/go/global"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func _Encode(entity interface{}, encodingType EncodingType) (string, error) {

	switch e := entity.(type) {
	case proto.Message:
		opts := protojson.MarshalOptions{
			EmitUnpopulated: true,
		}
		b, err := opts.Marshal(e)
		if err != nil {
			return "", err
		}

		switch encodingType {
		case EncodingType_BASE64:
			return _encodeToBase64(e)
		default:
			return string(b), nil
		}
	default:
		switch encodingType {
		case EncodingType_BASE64:
			var buf bytes.Buffer
			encoder := base64.NewEncoder(base64.StdEncoding, &buf)
			if err := json.NewEncoder(encoder).Encode(e); err != nil {
				return "", err
			}
			encoder.Close()
			return buf.String(), nil
		default:
			return "", nil
		}

	}
}

func _encodeToBase64(e interface{}) (str string, err error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	defer encoder.Close()
	if err = json.NewEncoder(encoder).Encode(e); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func DecodeFiberRequest(c *fiber.Ctx, out proto.Message) error {
	un := protojson.UnmarshalOptions{
		DiscardUnknown: true,
		AllowPartial:   true,
	}
	if err := un.Unmarshal(c.Body(), out); err != nil {
		return err
	}
	return nil

}
func _Decode(in string, out interface{}, encodingType EncodingType) error {
	switch encodingType {
	case EncodingType_BASE64:
		switch o := out.(type) {
		case proto.Message:
			un := protojson.UnmarshalOptions{
				DiscardUnknown: true,
				AllowPartial:   true,
			}
			decoded, err := base64.StdEncoding.DecodeString(in)
			if err != nil {
				if err = un.Unmarshal([]byte(in), o); err != nil {
					return err
				}
				return nil
			}

			if err = un.Unmarshal(decoded, o); err != nil {
				return err
			}
			return nil
		default:
			return json.NewDecoder(base64.NewDecoder(base64.StdEncoding,
				strings.NewReader(in))).Decode(out)
		}
	case EncodingType_JSON:
		if err := json.Unmarshal([]byte(in), out); err != nil {
			return fmt.Errorf("error unmarshalling json string to struct %v", err)
		}
		return nil
	default:
		return fmt.Errorf("Unrecognized encoding type of %d", encodingType)
	}

}

type Prefix string

const (
	None          Prefix = ""
	Guest         Prefix = "gst_"
	Guardian      Prefix = "grd_"
	Partner       Prefix = "ptn_"
	Journey       Prefix = "jny_"
	Operations    Prefix = "ops_"
	Experience    Prefix = "exp_"
	PaymentMethod Prefix = "pym_"
	Finance       Prefix = "fin_"
)

func NewULID(prefix Prefix) (string, error) {
	var final strings.Builder
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	if err != nil {
		return "", err
	}
	if prefix != None {
		final.WriteString(string(prefix))
		final.WriteString(id.String())
	}

	return final.String(), nil
}

func ValidateQuery(query *global.Query) error {
	if query == nil {
		return fmt.Errorf("the query was nil")
	} else if query.Terms == nil {
		return fmt.Errorf("no terms were provided when attempting to get rates")
	}
	return nil
}
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(stream); err != nil {
		return nil
	}
	return buf.Bytes()
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

func ConvertStructToMap(input interface{}, toSnake bool) map[string]interface{} {
	// Create an empty map to hold the converted data
	result := make(map[string]interface{})

	// Get the type and value of the input
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	// Ensure the input is a struct
	if inputType.Kind() != reflect.Struct {
		panic("Input must be a struct")
	}

	// Iterate through each field in the struct
	for i := 0; i < inputType.NumField(); i++ {
		// Get the field name and value
		fieldName := inputType.Field(i).Name
		fieldValue := inputValue.Field(i).Interface()
		if toSnake {
			fieldName = ToSnakeCase(fieldName)
		}
		// Assign the field name and value to the result map
		result[fieldName] = fieldValue
	}

	return result
}

func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result.WriteByte('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
