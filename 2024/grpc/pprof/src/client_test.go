package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	Hello "pprof/src/pb/this_is_go_package_content"
	"testing"
)

func TestPprofMarshal(t *testing.T) {
	people := &Hello.People{
		Name: "Anthony_4926",
		Age:  20,
	}

	hello := &Hello.HelloRequest{
		Who: people,
		Msg: "Hello, world! 2024/03/16 16:50:00",
	}

	marshalHello, err := proto.Marshal(hello)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("marshalHello: %s\n", string(marshalHello))
	}
}

func getPprofMarshal() []byte {
	people := &Hello.People{
		Name: "Anthony_4926",
		Age:  20,
	}

	hello := &Hello.HelloRequest{
		Who: people,
		Msg: "Hello, world! 2024/03/16 16:50:00",
	}

	marshalHello, _ := proto.Marshal(hello)
	return marshalHello

}

func TestPprofUnmarshal(t *testing.T) {
	s := getPprofMarshal()
	unmarshalHello := &Hello.HelloRequest{}
	err := proto.Unmarshal(s, unmarshalHello)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("unmarshalHello: %+v\n", unmarshalHello)
	}

}
