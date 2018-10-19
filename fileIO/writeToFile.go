package fileIO

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func WriteInFile(fileName string, pb proto.Message) error {
	out, err := proto.Marshal(pb) // marsahlling means serializing it in byte array.
	if err != nil {
		log.Fatalln("Error while marshalling...", err)
		return err
	}
	err = ioutil.WriteFile(fileName, out, 0644)
	if err != nil {
		return err
	}
	fmt.Println("File has been written in ", fileName)
	return nil
}