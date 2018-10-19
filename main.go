package main

import (
	"fmt"
	"log"
	"protobuffGolang/fileIO"
	"protobuffGolang/protoAndJson"
	"protobuffGolang/test1"
)

func main() {
	enumExample()
}
func jsonPbDemo() {
	sm := &test1.SimpleProto{
		Age: 182,
		Name: "Lucene",
		Address: "Pluto",
	}

	out, err := protoAndJson.ProtoToJson(sm)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out)
	// now convert this out back to proto to test if this really works.
	sm1 := &test1.SimpleProto{}
	protoAndJson.JsonToProto(out, sm1)
	fmt.Println("Json to proto: ", sm1)
}

func fileIODemo() {
	sm := simpleExample()
	err := fileIO.WriteInFile("simple.bin", sm)
	if err != nil {
		log.Fatal("Error")
	}

	readValue := &test1.SimpleProto{}
	fileIO.ReadFromFile("simple.bin", readValue)
	fmt.Println("Read from file: ", readValue)
}

func simpleExample() *test1.SimpleProto {
	sm := &test1.SimpleProto{
		Age: 12,
		Name: "Yoda",
		Address: "GolangVerser",
	}

	fmt.Println("Hello World: sm: ", sm)
	sm.Age = 45
	fmt.Println(sm)
	fmt.Println(sm.GetAddress())
	return sm
}

func enumExample() {
	enumProto := test1.EnumProto{}
	enumProto.Id = 76
	enumProto.WeekOfday = test1.DayOfWeek_SATURDAY

	fmt.Println("enumProto example: ", enumProto)

	enumProto2 := test1.EnumProto{
		Id: 89,
		WeekOfday:test1.DayOfWeek_FRIDAY,
	}

	fmt.Println("EnumProto2 Example: ", enumProto2)
}