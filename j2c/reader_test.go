package json2chan

import (
	"encoding/json"
	"testing"
)

var bytes = []byte(`{
  "list": [
    {"key":  "value1"},
    {"key":  "value2"},
    {"key":  "value3"},
    {"key":  "value4"}
  ]
}`)

func TestReadJsonObjects(t *testing.T) {
	expected := make(map[interface{}]string)
	expected["value1"] = ""
	expected["value2"] = ""
	expected["value3"] = ""
	expected["value4"] = ""

	in := make(chan byte)
	out := ReadObjects(in, "list")

	go func() {
		for _, b := range bytes {
			in <- b
		}
	}()

	var item = Item{}
	for s := range out {
		err := json.Unmarshal([]byte(s), &item)
		if err != nil {
			t.Error(err)
		}
		delete(expected, item.Key)
	}

	l := len(expected)
	if l > 0 {
		t.Fatal("Expected 0 items left over, not ", l)
	}

}

type Item struct {
	Key string `json:"key"`
}
