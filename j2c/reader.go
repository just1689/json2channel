package j2c

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

}
