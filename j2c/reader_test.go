package j2c

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

var bytesMessy = []byte(`{
  "type": "list",
  "list": [
    {"key":  "value1"},
    {"key":  "value2"},
    {"key":  "value3"},
    {"key":  "value4"}
  ]
}`)

var bytesNestedMessy = []byte(`{
  "type": {
	"name": "x",
    "list": [{"k": "v"}]
  },
  "list": [
    {"key":  "value1"},
    {"key":  "value2"},
    {"key":  "value3"},
    {"key":  "value4"}
  ]
}`)

func TestReadJsonObjects(t *testing.T) {
	expected := getExpectedMap()

	in := StartByteReader(bytes)
	out := ReadObjects(in, "list")

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

func TestReadJsonObjectsMessyNames(t *testing.T) {
	expected := getExpectedMap()

	in := StartByteReader(bytesMessy)
	out := ReadObjects(in, "list")

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

//func TestReadJsonObjectsNestedMessy(t *testing.T) {
//	expected := getExpectedMap()
//
//	in := StartByteReader(bytesNestedMessy)
//	out := ReadObjects(in, "list")
//
//	var item = Item{}
//	for s := range out {
//		err := json.Unmarshal([]byte(s), &item)
//		if err != nil {
//			t.Error(err)
//		}
//		delete(expected, item.Key)
//	}
//
//	l := len(expected)
//	if l > 0 {
//		t.Fatal("Expected 0 items left over, not ", l)
//	}
//
//}

type Item struct {
	Key string `json:"key"`
}

func getExpectedMap() map[interface{}]string {
	expected := make(map[interface{}]string)
	expected["value1"] = ""
	expected["value2"] = ""
	expected["value3"] = ""
	expected["value4"] = ""
	return expected
}
