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

	switch b {
	case LS:
		i.ls++
		break
	case RS:
		i.rs++
		break
	case LB:
		i.lb++
		break
	case RB:
		i.rb++
		break
	}

	return
}

func (i *Interpreter) Depth() int {
	return i.ls - i.rs
}

func (i *Interpreter) InArr() bool {
	return i.lb > i.rb
}
