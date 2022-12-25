package main

type Stack struct {
	crates []string
}

func NewStack() Stack {
	return Stack{make([]string, 0)}
}

func (stack *Stack) Pop(index int) (crate string) {
	crate = stack.crates[index]
	if index+1 == len(stack.crates) {
		stack.crates = stack.crates[:index]
	} else {
		stack.crates = append(stack.crates[:index], stack.crates[index+1:]...)
	}
	return
}

func (stack *Stack) PopLast() string {
	return stack.Pop(len(stack.crates) - 1)
}

func (stack *Stack) Append(crate string) {
	stack.crates = append(stack.crates, crate)
}
