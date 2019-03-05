package j2c

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func ReadObjectsToFunc(in chan byte, element string, f func(s string)) {
	out := ReadObjects(in, element)
	for o := range out {
		f(o)
	}
}

func ReadObjects(in chan byte, element string) (out chan string) {
	out = make(chan string)
	go func() {
		found := ReadUntilGivenElement(in, element)
		if !found {
			logrus.Error(errors.New(fmt.Sprint("Could not find", element, "in channel")))
			close(out)
			return
		}
		innerCh := ReadInnerArr(in)
		ReadJsonObjects(innerCh, out)
	}()
	return out
}

func ReadUntilGivenElement(in chan byte, element string) (found bool) {
	found = false
	var eof bool
	word := ""
	for !found {
		found, word, eof = ReadUntilNextElement(in)
		if eof { // ???
			found = false
			return
		}
		if found && word == element {
			found = true
			return
		}
	}
	return
}

func ReadUntilNextElement(in chan byte) (found bool, word string, eof bool) {
	isWord := false
	word = ""
	for letter := range in {
		if letter == QUOTE {
			if isWord {
				found = true
				return
			}
			isWord = true
			continue
		}
		if isWord {
			word += string(letter)
		}

	}
	return
}

func ReadInnerArr(in chan byte) (out chan byte) {
	out = make(chan byte)
	go func() {
		countLB := 0
		countRB := 0
		for c := range in {
			if c == LB {
				countLB++
			} else if c == RB {
				countRB++
				if countLB == countRB {
					close(out)
					return
				}
			}
			if countLB > 0 {
				out <- c
			}
		}
	}()
	return out
}

//ReadJsonObjects writes json objects to a channel
func ReadJsonObjects(in chan byte, out chan string) {
	go func() {
		object := ""
		countLS := 0
		countRS := 0
		reading := false
		for c := range in {
			if c == LS {
				reading = true
				countLS++
			} else if c == RS {
				countRS++
				if countLS == countRS {
					object += string(c)
					out <- object
					object = ""
					reading = false
					continue
				}
			} else if c == COMMA {
				if countLS == countRS {
					continue
				}
			}
			if countLS > 0 && reading {
				object += string(c)
			}
		}
		close(out)
		return

	}()
}
