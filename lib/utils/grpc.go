package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	// "go.starlark.net/lib/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type EncodingType int32

const (
	EncodingType_UNSPECIFIED EncodingType = iota
	EncodingType_BASE64
	EncodingType_BASE64JSON
	EncodingType_JSON
)

func EncodeV2(entity interface{}, encodingType EncodingType) ([]byte, error) {
	switch e := entity.(type) {
	case proto.Message:
		opts := protojson.MarshalOptions{
			EmitUnpopulated: true,
		}
		b, err := opts.Marshal(e)
		if err != nil {
			return nil, err
		}
		switch encodingType {
		case EncodingType_BASE64:
			return encodeToBase64(e)
		case EncodingType_JSON:
			return b, nil
		default:
			return nil, fmt.Errorf("unrecognized encoding type of %d", encodingType)
		}

	default:
		switch encodingType {
		case EncodingType_BASE64:
			var buf bytes.Buffer
			encoder := base64.NewEncoder(base64.StdEncoding, &buf)
			if err := json.NewEncoder(encoder).Encode(e); err != nil {
				return nil, err
			}
			encoder.Close()
			return buf.Bytes(), nil
		case EncodingType_JSON:
			return json.Marshal(e)

		default:
			return nil, fmt.Errorf("unrecognized encoding type of %d", encodingType)
		}

	}
}

func DecodeV2(in interface{}, out interface{}, encodingType EncodingType) error {
	// fmt.Println("about to decode the world ok")
	if in == nil || out == nil {
		return fmt.Errorf("both in and out must be non-nil")
	}

	switch in := in.(type) {
	case string:
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
					fmt.Printf("not base64 string %v, will try proto.Unmarshal", err)
					return err
				}

				if err = un.Unmarshal(decoded, o); err != nil {
					return err
				}

				return nil
			default:
				return fmt.Errorf("unrecognized type of %T", out)
				// return json.NewDecoder(base64.NewDecoder(base64.StdEncoding,
				// 	strings.NewReader(in))).Decode(out)
			}
		case EncodingType_JSON:
			switch o := out.(type) {
			case proto.Message:
				un := protojson.UnmarshalOptions{
					DiscardUnknown: true,
					AllowPartial:   true,
				}

				if err := un.Unmarshal([]byte(in), o); err != nil {
					return err
				}
			default:
				if err := json.Unmarshal([]byte(in), out); err != nil {
					return fmt.Errorf("error unmarshalling json string to struct %v", err)
				}
				return nil

			}
		default:
			return fmt.Errorf("unrecognized encoding type of %d", encodingType)
		}
	default:
		return fmt.Errorf("unrecognized type of %T", in)
	}
	return fmt.Errorf(" should not be! something was missed. this %T", in)

}

// TODO: fix this method so that it returns the correct encoding and not
// the bs base64 encoding that it is currently retuning.
func Encode(entity interface{}, encodingType EncodingType) ([]byte, error) {
	switch e := entity.(type) {
	case proto.Message:
		opts := protojson.MarshalOptions{
			EmitUnpopulated: true,
		}
		b, err := opts.Marshal(e)
		if err != nil {
			return nil, err
		}

		switch encodingType {
		case EncodingType_BASE64:
			return encodeToBase64(e)
		default:
			return b, nil
		}
	default:
		switch encodingType {
		case EncodingType_BASE64:
			var buf bytes.Buffer
			encoder := base64.NewEncoder(base64.StdEncoding, &buf)
			if err := json.NewEncoder(encoder).Encode(e); err != nil {
				return nil, err
			}
			encoder.Close()
			return buf.Bytes(), nil
		case EncodingType_JSON:
			return json.Marshal(e)

		default:
			return nil, fmt.Errorf("unrecognized encoding type of %d", encodingType)
		}

	}
}

func EncodeToBase64Str(entity interface{}) (string, error) {
	switch e := entity.(type) {
	case proto.Message:
		opts := protojson.MarshalOptions{
			EmitUnpopulated: true,
		}
		b, err := opts.Marshal(e)
		if err != nil {
			return "", err
		}

		return base64.StdEncoding.EncodeToString(b), nil

	default:

		return "", fmt.Errorf("unrecognized encoding")
	}

}

func encodeToBase64(e interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	defer encoder.Close()
	if err = json.NewEncoder(encoder).Encode(e); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

	// return []byte(base64.StdEncoding.EncodeToString(b)), nil

}

func Decode(in string, out interface{}, encodingType EncodingType) error {

	switch encodingType {
	case EncodingType_BASE64:
		// fmt.Println("got base64 encode option")

		switch o := out.(type) {
		case proto.Message:
			// fmt.Println("its a proto message")

			un := protojson.UnmarshalOptions{
				DiscardUnknown: true,
				AllowPartial:   true,
			}

			decoded, err := base64.StdEncoding.DecodeString(in)
			if err != nil {
				// fmt.Printf("error decoding base64 string %v", err)
				if err = un.Unmarshal([]byte(in), o); err != nil {
					// fmt.Printf("error decoding byte slice : %v\n", err)
					return err
				}

				// fmt.Println("decodd it i guess ")
				return nil
			}

			if err = un.Unmarshal(decoded, o); err != nil {
				// fmt.Printf("error unmashaling decoded into msg : %v\n", err)
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
		return fmt.Errorf("unrecognized encoding type of %d", encodingType)
	}

}

// 	switch encodingType {
// 	case EncodingType_BASE64:
// 		decoded, err := base64.StdEncoding.DecodeString(in)
// 		if err != nil {
// 			fmt.Printf("not base64 string %v, will try proto.Unmarshal", err)
// 			return err
// 		}
// 		return un.Unmarshal(decoded, o)
// 	case EncodingType_JSON:
// 		if err := un.Unmarshal([]byte(in), o); err != nil {
// 			fmt.Printf("error decoding json encoding : %v\n", err)
// 			return err
// 		}
// 	default:
// 		return fmt.Errorf("unrecognized encoding type of %d", encodingType)
// 	}

// default:
// 	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding,
// 		strings.NewReader(in))).Decode(out)
