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

var bytesRootless = []byte(`
[
   {
      "entity":"car",
      "id":"1",
      "v":{
         "name":"justin"
      }
   },
   {
      "entity":"car",
      "id":"2",
      "v":{
         "name":"justin"
      }
   },
   {
      "entity":"user",
      "id":"5",
      "v":{
         "x":"y",
         "body":{
            "arr":[
               {
                  "id":1
               },
               {
                  "id":2
               },
               {
                  "id":3
               }
            ]
         },
         "Whaaat":"ytou never"
      }
   },
   {
      "entity":"user",
      "id":"0",
      "v":{
         "x":"y",
         "body":{
            "arr":[
               {
                  "id":1
               },
               {
                  "id":2
               },
               {
                  "id":3
               }
            ]
         },
         "Whaaat":"ytou never"
      }
   },
   {
      "entity":"user",
      "id":"11",
      "v":{
         "x":"y",
         "body":{
            "arr":[
               {
                  "id":1
               },
               {
                  "id":2
               },
               {
                  "id":3
               }
            ]
         },
         "Whaaat":"ytou never"
      }
   }
]`)

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

func TestReadJsonObjectsRootless(t *testing.T) {

	in := StartByteReader(bytesRootless)
	out := ReadObjects(in, ".")

	count := 0
	for _ = range out {
		count++
	}
	if count != 5 {
		t.Error("Intead of 5 rows returned, returned ", count)
		return
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
