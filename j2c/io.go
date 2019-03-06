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
		f.len = len(f.current.Bytes)
	}

	if f.position >= f.len {
		eof = true
		return
	}

	b = f.current.Bytes[f.position]
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
				Bytes: make([]byte, 4096),
			}

			p.Bytes = p.Bytes[:cap(p.Bytes)]
			n, err := f.Read(p.Bytes)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}
			p.Bytes = p.Bytes[:n]
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
			Bytes: bytes,
		}
		fileReader.in <- &frame{
			eof: false,
		}
	}()
	return fileReader
}

type frame struct {
	Bytes []byte
	eof   bool
}
