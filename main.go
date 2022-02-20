package main

import (
	"awesomeProject/serialization/models"
	byteslib "bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hamba/avro"
	"github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type message struct {
	ID        int            `json:"id" avro:"id"`
	StrData   string         `json:"strData" avro:"strData"`
	FloatData float32        `json:"floatData" avro:"floatData"`
	DataArray []string       `json:"dataArray" avro:"dataArray"`
	DataMap   map[string]int `json:"dataMap" avro:"dataMap"`
}

type messageXml struct {
	ID        int            `json:"id" avro:"id"`
	StrData   string         `json:"strData" avro:"strData"`
	FloatData float32        `json:"floatData" avro:"floatData"`
	DataArray []string       `json:"dataArray" avro:"dataArray"`
}

func main() {

	mapTest := make(map[string]int)
	mapTest["test1"] = 1
	mapTest["test2"] = 2
	msg := message{
		ID:        1,
		StrData:   "test",
		FloatData: 1.023,
		DataArray: []string{"132", "test", "rsdf"},
		DataMap:   mapTest,
	}

	var network byteslib.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	start := time.Now()
	err := enc.Encode(msg)
	if err != nil {
		log.Fatal(err)
	}
	durationSer := time.Since(start)
	netBytes := network.Bytes()
	var msgnet message
	start = time.Now()
	err = dec.Decode(&msgnet)
	if err != nil {
		log.Fatal(err)
	}
	durationDes := time.Since(start)

	fmt.Println("Go serialization")
	fmt.Println("Time serialization: ", durationSer)
	fmt.Println("Time deserialization: ", durationDes)
	fmt.Println("Length string: ", len(string(netBytes)))

	//json
	start = time.Now()
	bytesJson, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	durationSer = time.Since(start)

	msg2 := message{}
	start = time.Now()
	err = json.Unmarshal(bytesJson, &msg2)
	if err != nil {
		log.Fatal(err)
	}
	durationDes = time.Since(start)
	fmt.Println()
	fmt.Println("Json serialization")
	fmt.Println("Time serialization: ", durationSer)
	fmt.Println("Time deserialization: ", durationDes)
	fmt.Println("Length string: ", len(string(bytesJson)))
	//json

	//yaml
	start = time.Now()
	bytesYaml, err := yaml.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	durationSer = time.Since(start)

	msgYaml := message{}
	start = time.Now()
	err = yaml.Unmarshal(bytesYaml, &msgYaml)
	if err != nil {
		log.Fatal(err)
	}
	durationDes = time.Since(start)
	fmt.Println()
	fmt.Println("Yaml serialization")
	fmt.Println("Time serialization: ", durationSer)
	fmt.Println("Time deserialization: ", durationDes)
	fmt.Println("Length string: ", len(string(bytesYaml)))
	//yaml

	//yaml
	start = time.Now()
	bytesMsgPack, err := msgpack.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	durationSer = time.Since(start)

	msgPack := message{}
	start = time.Now()
	err = msgpack.Unmarshal(bytesMsgPack, &msgPack)
	if err != nil {
		log.Fatal(err)
	}

	durationDes = time.Since(start)
	fmt.Println()
	fmt.Println("MsgPack serialization")
	fmt.Println("Time serialization: ", durationSer)
	fmt.Println("Time deserialization: ", durationDes)
	fmt.Println("Length string: ", len(string(bytesMsgPack)))
	//yaml

	//avro

	schemaStr, err := ioutil.ReadFile("schema.avsc")
	if err != nil {
		log.Fatal(err)
	}

	schema, err := avro.Parse(string(schemaStr))

	start = time.Now()
	bytesAvro, err := avro.Marshal(schema, msg)
	if err != nil {
		log.Fatal(err)
	}
	durationSer = time.Since(start)

	msgAvro := message{}
	start = time.Now()
	err = avro.Unmarshal(schema, bytesAvro, &msgAvro)
	if err != nil {
		log.Fatal(err)
	}

	durationDes = time.Since(start)
	fmt.Println()
	fmt.Println("Avro serialization")
	fmt.Println("Time serialization: ", durationSer)
	fmt.Println("Time deserialization: ", durationDes)
	fmt.Println("Length string: ", len(string(bytesAvro)))
	//avro

	//protobuf
	mapProto := make(map[string]int32)
	mapProto["test1"] = 1
	mapProto["test2"] = 2

	start = time.Now()
	out, err := proto.Marshal(&models.Message{
		Id:        int32(msg.ID),
		StrData:   msg.StrData,
		DataMap:   mapProto,
		DataArray: msg.DataArray,
	})
	durationSer = time.Since(start)
	if err != nil {
		log.Fatal(err)
	}
	msgProto := models.Message{}
	start = time.Now()
	err = proto.Unmarshal(out, &msgProto)
	if err != nil {
		log.Fatal(err)
	}

	durationDes = time.Since(start)
	fmt.Println()
	fmt.Println("Protobuf serialization")
	fmt.Println("Time serialization: ", durationSer)
	fmt.Println("Time deserialization: ", durationDes)
	fmt.Println("Length string: ", len(string(out)))
	//protobuf

	//xml
	//map xml не поддерживает
	msgXmlTest := messageXml{
		ID:        1,
		StrData:   "test",
		FloatData: 1.023,
		DataArray: []string{"132", "test", "rsdf"},
	}

	start = time.Now()
	bytesXml, err := xml.Marshal(msgXmlTest)
	if err != nil {
		log.Fatal(err)
	}
	durationSer = time.Since(start)


	msgXml := messageXml{}
	start = time.Now()
	err = xml.Unmarshal(bytesXml, &msgXml)
	if err != nil {
		log.Fatal(err)
	}

	durationDes = time.Since(start)
	fmt.Println()
	fmt.Println("XML serialization")
	fmt.Println("Time serialization: ", durationSer)
	fmt.Println("Time deserialization: ", durationDes)
	fmt.Println("Length string: ", len(string(bytesMsgPack)))
	//xml
}
