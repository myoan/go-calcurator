package main

import (
	"fmt"
	"github.com/myoan/go-calcurator/stack"
)

type CalcType int

const (
	Int CalcType = iota
	Add
	Sub
	Mul
	Div
	Eof
)

type Data struct {
	tipe  CalcType
	value int
}

func Calc(data []Data) int {
	reg := stack.CalcStack{Index: 0}
	for _, d := range data {
		switch d.tipe {
		case Int:
			reg.Push(d.value)
		case Add:
			a := reg.Pop()
			b := reg.Pop()
			reg.Push(b + a)
		case Sub:
			a := reg.Pop()
			b := reg.Pop()
			reg.Push(b - a)
		case Mul:
			a := reg.Pop()
			b := reg.Pop()
			reg.Push(b * a)
		case Div:
			a := reg.Pop()
			b := reg.Pop()
			reg.Push(b / a)
		case Eof:
			break
		}
	}
	return reg.Result()
}

func main() {
	/*
		data := []Data{
			{tipe: Int, value: 5},
			{tipe: Int, value: 2},
			{tipe: Add, value: 1},
		}
	*/

	// (5 + 2) * 3
	/*
		data := []Data{
			{tipe: Int, value: 5},
			{tipe: Int, value: 2},
			{tipe: Add, value: 1},
			{tipe: Int, value: 3},
			{tipe: Mul, value: 1},
			{tipe: Eof, value: 1},
		}
	*/
	// (5 + 2) * (3 - 1)
	data := []Data{
		{tipe: Int, value: 5},
		{tipe: Int, value: 2},
		{tipe: Add, value: 1},
		{tipe: Int, value: 3},
		{tipe: Int, value: 1},
		{tipe: Sub, value: 1},
		{tipe: Mul, value: 1},
		{tipe: Eof, value: 1},
	}

	fmt.Println(Calc(data))
}
