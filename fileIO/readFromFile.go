package fileIO

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func ReadFromFile(fileName string, pb proto.Message) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err = proto.Unmarshal(data, pb); err!= nil {
		return err
	}
	return nil
}
