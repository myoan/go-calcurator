package main

import "fmt"

type CalcType int

const (
	Int CalcType = iota
	Opr
	Eof
)

type CalcOpr int

const (
	Add CalcOpr = iota
	Sub
	Mul
	Div
)

type Data struct {
	tipe  CalcType
	value int
}

func Calc(data []Data) int {
	reg := [10]int{}
	i := 0
	for _, d := range data {
		switch d.tipe {
		case Int:
			reg[i] = d.value
			i = i + 1
		case Opr:
			a := reg[i-2]
			b := reg[i-1]
			reg[i] = a + b
		case Eof:
			break
		}
	}
	return reg[i]
}

func main() {
	data := []Data{
		{tipe: Int, value: 5},
		{tipe: Int, value: 1},
		{tipe: Opr, value: 1},
		{tipe: Eof, value: 1},
	}

	fmt.Println(Calc(data))
}
