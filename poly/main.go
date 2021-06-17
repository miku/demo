// https://i.imgur.com/NSgBYYM.png
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

// We have different types of messages.
type Type string

const (
	start Type = "start"
	stop  Type = "stop"
	mark  Type = "mark"
)

// Start, Stop and Mark are different message types with different fields.

type Start struct {
	Pos int `json:"pos"`
}

type Stop struct {
	Pos    int    `json:"pos"`
	Reason string `json:"reason"`
}

type Mark struct {
	Pos   int `json:"pos"`
	Color int `json:"color"`
}

// Message is an envelope; data can hold any of Start, Stop, Mark (and others).

type Message struct {
	ID   string      `json:"id"`
	T    string      `json:"t"`
	Type Type        `json:"type"`
	Data interface{} `json:"data"`
}

// String just customizes the output.
func (m *Message) String() string {
	var (
		yellow = color.New(color.FgYellow).SprintFunc()
		red    = color.New(color.FgRed).SprintFunc()
		green  = color.New(color.FgGreen).SprintFunc()
		dot    string
	)
	switch m.Type {
	case start:
		dot = green("■")
	case stop:
		dot = red("■")
	case mark:
		dot = yellow("■")
	default:
		dot = "■"
	}
	return fmt.Sprintf("[%s % 6s] %s %s %+v", dot, m.Type, m.ID, m.T, m.Data)
}

// UnmarshalJSON will create a message find out the right type for the Data field by dummy unmarshaling.
func (m *Message) UnmarshalJSON(p []byte) (err error) {
	var dummy struct {
		ID   string          `json:"id"`
		T    string          `json:"t"`
		Type Type            `json:"type"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(p, &dummy); err != nil {
		return err
	}
	m.ID = dummy.ID
	m.T = dummy.T
	m.Type = dummy.Type
	switch dummy.Type {
	case start:
		v := Start{}
		err = json.Unmarshal(dummy.Data, &v)
		m.Data = v
	case stop:
		v := Stop{}
		err = json.Unmarshal(dummy.Data, &v)
		m.Data = v
	case mark:
		v := Mark{}
		err = json.Unmarshal(dummy.Data, &v)
		m.Data = v
	default:
		return fmt.Errorf("unknown message type: %v", dummy.Type)
	}
	return
}

// readStream reads messages and prints them the to stdout.
func readStream(r io.Reader) error {
	br := bufio.NewReader(r)
	for {
		b, err := br.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var msg Message
		if err := json.Unmarshal(b, &msg); err != nil {
			log.Fatal(err)
		}
		fmt.Println(&msg)
	}
	return nil
}

func main() {
	if err := readStream(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
