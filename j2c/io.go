package j2c

import (
	"fmt"
	"io"
	"os"
)

type ChanReader struct {
	in       chan *frame
	current  *frame
	position int
	len      int
}

type Reader interface {
	Next() (b byte, eof bool)
}

func (f *ChanReader) Next() (b byte, eof bool) {

	if f.current == nil || f.position >= f.len {
		item := <-f.in
		if item.eof {
			eof = true
			return
		}
		f.current = item
		f.position = 0
		f.len = len(f.current.bytes)
	}

	if f.position >= f.len {
		eof = true
		return
	}

	b = f.current.bytes[f.position]
	f.position++

	return
}

func StartFileReader(filename string) *ChanReader {

	fileReader := &ChanReader{}
	fileReader.in = make(chan *frame, 256)
	go func() {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		for {
			p := &frame{
				bytes: make([]byte, 4096),
			}

			p.bytes = p.bytes[:cap(p.bytes)]
			n, err := f.Read(p.bytes)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}
			p.bytes = p.bytes[:n]
			fileReader.in <- p

		}

		p := &frame{
			eof: false,
		}
		fileReader.in <- p

	}()
	return fileReader
}

func StartByteReader(bytes []byte) *ChanReader {

	fileReader := &ChanReader{}
	fileReader.in = make(chan *frame, 1)

	go func() {
		fileReader.in <- &frame{
			bytes: bytes,
		}
		fileReader.in <- &frame{
			eof: false,
		}
	}()
	return fileReader
}

type frame struct {
	bytes []byte
	eof   bool
}
