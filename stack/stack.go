package stack

type CalcStack struct {
	reg   [10]int
	Index int
}

func (st *CalcStack) Result() int {
	return st.reg[st.Index-1]
}

func (st *CalcStack) Pop() int {
	result := st.reg[st.Index-1]
	st.Index = st.Index - 1
	return result
}

func (st *CalcStack) Push(r int) {
	st.reg[st.Index] = r
	st.Index = st.Index + 1
}
