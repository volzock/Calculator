package structs

import (
	"errors"
)

type stackNode struct {
	before *stackNode
	value  string
}

type Stack struct {
	last *stackNode
	size int
}

func (stack *Stack) Add(s *string) {
	if stack.last == nil {
		stack.last = new(stackNode)
		stack.last.value = *s
		stack.size = 1
	} else {
		var newNode *stackNode = new(stackNode)
		newNode.value = *s
		newNode.before = stack.last
		stack.last = newNode
		stack.size += 1
	}
}

func (stack *Stack) Pop() (v string, e error) {
	if stack.size == 0 {
		e = errors.New("stack is emplty")
	} else {
		stack.size -= 1
		v = stack.last.value
		stack.last = stack.last.before
	}

	return v, e
}
