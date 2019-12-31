package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/worldofprasanna/protobuf/messages"
	"github.com/golang/protobuf/proto"
	"encoding/binary"
)

func main() {
	flag.Parse()
	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		msg := strings.Join(flag.Args()[1:], " ")
		err = add(msg)
	default:
		err = fmt.Errorf("invalid command %s", cmd)
	}

	if err != nil {
		fmt.Println("error occurred %v", err)
		os.Exit(1)
	}
}

func list() error {
	fmt.Println("list method invoked")
	return nil
}

const (
	protoFilePath = "protobuf.msg"
)

func add(msg string) error {
	fmt.Printf("add method invoked with %s", msg)
	m := &messages.Msg{
		Text: msg,
		To: "always@me",
	}
	msgBytes, err := proto.Marshal(m)
	if err != nil {
		return fmt.Errorf("error in marshalling the value %v", err)
	}

	f, err := os.OpenFile(protoFilePath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("error in opening the file %v", err)
	}

	if err = binary.Write(f, binary.LittleEndian, int64(len(msgBytes))); err != nil {
		return fmt.Errorf("error in writing length of message %v", err)
	}

	_, err = f.Write(msgBytes)
	if err != nil {
		return fmt.Errorf("error in writing the message %v", err)
	}

	if err = f.Close(); err != nil {
		return fmt.Errorf("error in closing the file %v", err)
	}

	return nil
}
