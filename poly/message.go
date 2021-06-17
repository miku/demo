package main

import (
	"encoding/json"
	"time"
)

// TODO: different messages plus an envelop or wrapper with extra information
// and a pointer to the message. Use interface first.
// https://making.pusher.com/alternatives-to-sum-types-in-go/
// https://eli.thegreenplace.net/2018/go-and-algebraic-data-types/
//
// Have []byte (json), want static fields, and a dynamic field, depending on
// some other fields (e.g. the type of message).

type Message struct {
	MID  string
	PID  string
	T    time.Time
	Type int
	Data interface{}
}

func (m *Message) FromBytes(p []byte, t interface{}) error {
	m.Data = t
	return json.Unmarshal(p, m)
}

func main() {
	example = `{"mid":"1","pid":"2","t":"123","type":"1","data":{"a":"1","b":"2"}}`
	msg := Message{}
	// ...
}
