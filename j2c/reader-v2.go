package j2c

import "fmt"

func ReadObjectsV2(in Reader, element string) (out chan string) {
	out = make(chan string)

	go func() {

		interpreter := BuildInterpreter(in)

		//Need new implementation of `ReadUntilGivenElement()`
		found := ReadUntilGivenElement(interpreter, element)
		if !found {
			close(out)
		}

		readArr(interpreter, out)
		close(out)

	}()
	return out

}

func readArr(interpreter *Interpreter, out chan string) {
	var b byte
	var eof bool
	word := ""
	started := false
	interpreter.ResetBrackets()
	for {
		b, eof = interpreter.Next()
		if eof {
			return
		}

		if interpreter.lb < 1 {
			continue
		}
		if b == LB && interpreter.lb == 1 {
			//Array starting
			interpreter.ResetSquiggly()
			continue
		}
		if interpreter.lb == interpreter.rb {
			return
		}

		if b == LS && interpreter.ls == 1 {
			//fmt.Println("started")
			started = true
		}

		//fmt.Println(string(b), interpreter.ls, interpreter.rs)

		if interpreter.ls == 0 {
			continue
		}

		if interpreter.ls == interpreter.rs && b == RS {

			if !started {
				fmt.Println("Not going to send ", word)
				continue
			}

			//fmt.Println("Writing to chan")
			word += string(b)
			out <- word
			word = ""
			interpreter.ResetSquiggly()
			started = false

			continue

		}

		word += string(b)

	}

}
