package main

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/json2channel/j2c"
	"time"
)

func main() {

	//runFileExample()
	runLargeExample()

}

func runFileExample() {
	in := j2c.StartFileReader("test.json")
	out := j2c.ReadObjects(in, "list")

	for o := range out {
		fmt.Println(o)
	}

}

func runLargeExample() {
	fmt.Println("Preparing for large example")
	silly := Wrapper{
		List: make([]Item, 10*1000*1000),
	}
	b, _ := json.Marshal(silly)
	fmt.Println("Prep done.")

	in := j2c.StartByteReader(b)
	out := j2c.ReadObjects(in, "list")

	fmt.Println("Starting...")
	start := time.Now()
	count := 0
	for range out {
		count++
	}
	duration := time.Since(start)
	ips := int(float64(count) / duration.Seconds())
	fmt.Println("Finished", count, duration, "@ ", ips, " items per second")

}

func makeGiantSlice() (w *Wrapper) {
	w = &Wrapper{
		List: make([]Item, 1000*1000),
	}
	for i, _ := range w.List {
		w.List[i] = Item{
			Key: "value",
		}
	}
	return
}

type Wrapper struct {
	List []Item `json:"list"`
}

type Item struct {
	Key string `json:"key"`
}
