package main

import (
	"bufio"
	"fmt"
	"github.com/myoan/go-calcurator/shunting_yard"
	"github.com/myoan/go-calcurator/stack"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CalcType int

const (
	Int CalcType = iota
	Add
	Sub
	Mul
	Div
	Pow
	Eof
)

type Data struct {
	tipe  CalcType
	value int
}

func ToData(token string) Data {
	numReg := regexp.MustCompile(`[0-9]+`)
	if numReg.Match([]byte(token)) {
		i, _ := strconv.Atoi(token)
		return Data{tipe: Int, value: i}
	} else {
		if strings.Compare("^", token) == 0 {
			return Data{tipe: Pow, value: 1}
		} else if strings.Compare("*", token) == 0 {
			return Data{tipe: Mul, value: 1}
		} else if strings.Compare("/", token) == 0 {
			return Data{tipe: Div, value: 1}
		} else if strings.Compare("+", token) == 0 {
			return Data{tipe: Add, value: 1}
		} else if strings.Compare("-", token) == 0 {
			return Data{tipe: Sub, value: 1}
		}
	}
	return Data{}
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
	var sc = bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		var str = ""
		if sc.Scan() {
			str = sc.Text()
		}
		if strings.Compare("q", str) == 0 {
			break
		}
		tokens := strings.Fields(str)
		rpnData := shunting_yard.ToRpn(tokens)
		data := []Data{}
		for _, d := range rpnData {
			data = append(data, ToData(d))
		}
		fmt.Println(Calc(data))
	}
}
