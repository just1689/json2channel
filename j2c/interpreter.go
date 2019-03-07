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

func (i *Interpreter) ResetBrackets() {
	i.lb = 0
	i.rb = 0
}

func (i *Interpreter) ResetSquiggly() {
	i.ls = 0
	i.rs = 0
}

func (i *Interpreter) Next() (b byte, eof bool) {
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

func (i *Interpreter) IsStarted() bool {
	return i.ls-i.rs == 1
}
