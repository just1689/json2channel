package j2c

func BuildInterpreter(r Reader) *Interpreter {
	return &Interpreter{
		reader: r,
	}
}

type Interpreter struct {
	reader         Reader
	ls, rs, lb, rb int
}

func (i *Interpreter) Read() (b byte, eof bool) {
	b, eof = i.reader.Next()
	if b == LS {
		i.ls++
	} else if b == RS {
		i.rs++
	} else if b == LB {
		i.lb++
	} else if b == RB {
		i.rb++
	}

	return
}

func (i *Interpreter) Depth() int {
	return i.ls - i.rs
}

func (i *Interpreter) InArr() bool {
	return i.lb > i.rb
}
