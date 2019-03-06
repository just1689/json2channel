# Json 2 Channel
Reads a json stream and writes items from the json array to a channel

Have you ever had to retrieve items in an array that are part of a large json payload? This might help!

# Usage

Given the string
```json
{
  "list": [
    {"key":  "value1"},
    {"key":  "value2"},
    {"key":  "value3"},
    {"key":  "value4"}
  ]
}
```

Call the library

`out := j2s.ReadObjects(ch, "list")`

- `ch` is a channel that accepts bytes.
- The output `out` is a channel that receives objects as strings as they are made available.

# Performance 
Currently, in small tests the library produces over 80k items per second.

 
