package redis

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/tonychill/ifitu/lib/utils"
	"google.golang.org/protobuf/proto"
)

type _journey struct {
	Id             string `json:"id"`
	RequestId      string `json:"request_id"`
	CreatorId      string `json:"creator_id"`
	CreatorType    string `json:"creator_type"`
	CreatorTypeStr string `json:"creator_type_str"`
	Name           string `json:"name"`
	StartDate      int64  `json:"start_date"`
	EndDate        int64  `json:"end_date"`
	Data           string `json:"data"`
}

func (d Document) ToProto(msg proto.Message) (proto.Message, error) {
	_jny := &_journey{}
	if err := utils.Decode(d.Properties["$"].(string), _jny, utils.EncodingType_JSON); err != nil {
		log.Error().Err(err).
			Msg("error decoding reponse from redis.")
		return nil, fmt.Errorf("error decoding repsonse")
	}

	if err := utils.Decode(_jny.Data, msg, utils.EncodingType_BASE64); err != nil {
		log.Error().Err(err).Msg("error decoding journey data.")
		return nil, err
	}

	return msg, nil
}

func (d Document) Decode(out interface{}) error {
	if err := utils.DecodeV2(d.Properties["$"].(string), out, utils.EncodingType_JSON); err != nil {
		log.Error().Err(err).
			Msgf("error decoding redis document with type %T", out)
		return fmt.Errorf("error decoding redis document")
	}

	return nil
}

// Deprecated: use doc.Decode instead
func (d Document) ToProtoV2(out proto.Message) error {
	if err := utils.DecodeV2(d.Properties["$"].(string), out, utils.EncodingType_JSON); err != nil {
		log.Error().Err(err).
			Msgf("error decoding redis document with proto type %T", out)
		return fmt.Errorf("error decoding redis document")
	}

	return nil
}

// func (d Document) loadFields(lst []interface{}) *Document {
// 	for i := 0; i < len(lst); i += 2 {
// 		var prop string
// 		switch lst[i].(type) {
// 		case []byte:
// 			prop = string(lst[i].([]byte))
// 		default:
// 			prop = lst[i].(string)
// 		}

// 		var val interface{}
// 		switch v := lst[i+1].(type) {
// 		case []byte:
// 			val = string(v)
// 		default:
// 			val = v
// 		}
// 		*d = d.Set(prop, val)
// 	}
// 	return d
// }
