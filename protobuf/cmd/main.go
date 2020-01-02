package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/worldofprasanna/protobuf/messages"
	"github.com/golang/protobuf/proto"
	"encoding/binary"
	"io/ioutil"
	"bytes"
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
		fmt.Printf("error occurred %v\n", err)
		os.Exit(1)
	}
}

func list() error {
	b, err := ioutil.ReadFile(protoFilePath)
	if err != nil {
		return fmt.Errorf("error in reading the file %v", err)
	}

	for {
		if len(b) == 0 {
			return nil
		} else if len(b) < sizeOfLength {
			return fmt.Errorf("invalid custom protobuf message with extra length %d", len(b))
		}

		var l length
		var err error
		if err = binary.Read(bytes.NewReader(b[:sizeOfLength]), binary.LittleEndian, &l); err != nil {
			return fmt.Errorf("error occurred when reading the length of message %v", err)
		}

		b = b[sizeOfLength:]

		var m messages.Msg
		fmt.Printf("Length of message %v\n", l)
		if err := proto.Unmarshal(b[:l], &m); err != nil {
			return fmt.Errorf("could not read the message %v", err)
		}
		b = b[l:]
		fmt.Printf(" %s\n", m.Text)
	}
}

type length int64

const (
	protoFilePath = "protobuf.msg"
	sizeOfLength = 8
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

	if err = binary.Write(f, binary.LittleEndian, length(len(msgBytes))); err != nil {
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
