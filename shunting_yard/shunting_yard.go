package shunting_yard

import (
	"fmt"
	"regexp"
	"strings"
)

type ParseStack struct {
	stack [10]string
	index int
}

func (ps *ParseStack) Result() string {
	return ps.stack[ps.index-1]
}

func (ps *ParseStack) IsEmpty() bool {
	return ps.index == 0
}

func (ps *ParseStack) Pop() string {
	result := ps.stack[ps.index-1]
	ps.index = ps.index - 1
	fmt.Println("Pop:", result)
	fmt.Println(ps)
	return result
}

func (ps *ParseStack) Push(r string) {
	ps.stack[ps.index] = r
	ps.index = ps.index + 1
	fmt.Println("Push:", r)
	fmt.Println(ps)
}

type ShuntingYard struct {
	numQueue []string
	oprStack ParseStack
}

func (sy *ShuntingYard) Add(d string) {
	fmt.Println("numQueue:", sy.numQueue)
	fmt.Println("oprStack:", sy.oprStack)
	numReg := regexp.MustCompile(`[0-9]+`)
	if numReg.Match([]byte(d)) {
		sy.numQueue = append(sy.numQueue, d)
	} else {
		for {
			if sy.oprStack.IsEmpty() {
				break
			} else if oprPriority(sy.oprStack.Result()) > oprPriority(d) {
				break
			}
			opr := sy.oprStack.Pop()
			sy.numQueue = append(sy.numQueue, opr)
		}
		sy.oprStack.Push(d)
	}
}

func oprPriority(d string) int {
	if strings.Compare("^", d) == 0 {
		return 1
	} else if strings.Compare("*", d) == 0 {
		return 2
	} else if strings.Compare("/", d) == 0 {
		return 2
	} else if strings.Compare("+", d) == 0 {
		return 3
	} else if strings.Compare("-", d) == 0 {
		return 3
	} else {
		return 1000
	}
}

func ToRpn(data []string) []string {
	sy := ShuntingYard{}
	for _, d := range data {
		sy.Add(d)
	}
	for {
		if sy.oprStack.IsEmpty() {
			break
		}
		d := sy.oprStack.Pop()
		sy.numQueue = append(sy.numQueue, d)
	}
	return sy.numQueue
}
