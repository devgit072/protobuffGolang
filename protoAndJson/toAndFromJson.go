package protoAndJson

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtoToJson(pb proto.Message) (string, error) {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)
	if err != nil {
		return "", err
	}
	fmt.Println("Jsonified : ", out)
	return out, nil
}

func JsonToProto(jsonValue string, pb proto.Message) error {
	err := jsonpb.UnmarshalString(jsonValue, pb)
	if err != nil {
		return err
	}
	return nil
}
