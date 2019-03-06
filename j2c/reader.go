package j2c

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func ReadObjects(in Reader, element string) (out chan string) {
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

func ReadUntilGivenElement(in Reader, element string) (found bool) {
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

func ReadUntilNextElement(in Reader) (found bool, word string, eof bool) {
	isWord := false
	word = ""
	var letter byte

	for {
		letter, eof = in.Next()
		if eof {
			return
		}
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

	eof = true
	return
}

func ReadInnerArr(in Reader) (out chan byte) {
	out = make(chan byte)
	go func() {
		countLB := 0
		countRB := 0
		var c byte
		var eof bool
		for {
			c, eof = in.Next()
			if eof {
				close(out)
				return
			}
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
