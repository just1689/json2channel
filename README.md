# Json 2 Channel
<img src="https://goreportcard.com/badge/github.com/just1689/json2channel" />&nbsp;<a href="https://codeclimate.com/github/just1689/json2channel/maintainability"><img src="https://api.codeclimate.com/v1/badges/d1d1fac9de319a350ca3/maintainability" /></a>
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
Currently, in small tests the library produces over 55,000 items per second.

 
