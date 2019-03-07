package j2c

func ReadObjects(in Reader, element string) (out chan string) {
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
			interpreter.ResetSquiggly()
			continue
		}
		if interpreter.lb == interpreter.rb {
			return
		}

		if interpreter.ls == 0 {
			continue
		}

		if interpreter.ls == interpreter.rs && b == RS {

			word += string(b)
			out <- word
			word = ""
			interpreter.ResetSquiggly()

			continue

		}

		word += string(b)

	}

}
