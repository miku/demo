package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Attribute struct {
	ID          int64           `json:"id"`
	Label       string          `json:"label"`
	MultiSelect bool            `json:"multiSelect"`
	RawValues   json.RawMessage `json:"values"`
	Values      []AttributeValue
}

func (a Attribute) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &a.RawValues); err != nil {
		return err
	}

	switch a.RawValues[0] {
	case '[':
		var v []AttributeValue
		err := json.Unmarshal(b, &v)
		if err != nil {
			return err
		}
		a.Values = v
		return nil
	case '{':
		var v AttributeValue
		err := json.Unmarshal(b, &v)
		if err != nil {
			return err
		}
		a.Values = []AttributeValue{v}
		return nil
	}

	return fmt.Errorf("unexpected attribute values: %T: %s", a.RawValues, a.RawValues)
}

type AttributeValue struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

func main() {
	dec := json.NewDecoder(os.Stdin)
	for {
		var attr Attribute
		err := dec.Decode(&attr)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(attr)
	}
}
