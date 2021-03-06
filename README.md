# Json 2 Channel
<img src="https://goreportcard.com/badge/github.com/just1689/json2channel" />&nbsp;<a href="https://codeclimate.com/github/just1689/json2channel/maintainability"><img src="https://api.codeclimate.com/v1/badges/d1d1fac9de319a350ca3/maintainability" /></a>&nbsp;<a href="https://codebeat.co/projects/github-com-just1689-json2channel-master"><img alt="codebeat badge" src="https://codebeat.co/badges/6e12c23a-78aa-4da7-a536-2081c91f298d" /></a>


Have you ever had to retrieve millions of items in a json array that multi gigabyte payload? This might help!

This library reads a json stream, picks out items from a json array and writes them to a channel.


# State of the project
<img src="https://img.shields.io/badge/State-unstable-dd1111.svg" />&nbsp;<img src="https://europe-west1-captains-badges.cloudfunctions.net/function-clone-badge-pc?project=just1689/json2channel" /><br />

# Usage

Given the file `test.json`
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

```go
in := j2c.StartFileReader("test.json")
out := j2c.ReadObjects(in, "list")

for o := range out {
    fmt.Println(o)
}

```


`in` is a instance whose struct implements Reader defined as
```go
type Reader interface {
	Next() (b byte, eof bool)
}
```

`out := j2c.ReadObjects(in, "list")` is the line that does the work.
The output `out` is a channel that receives objects as strings as they are made available.

Output

`{"key":  "value1"}`

`{"key":  "value2"}`

`{"key":  "value3"}`

`{"key":  "value4"}`



# Performance 
Currently, in small tests the library produces over 500k items per second.

 
