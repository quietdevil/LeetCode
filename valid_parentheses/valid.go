package main

import "fmt"

func main() {
	is := isValid("[]([])")
	fmt.Println(is)
	is = isValid("){")
	fmt.Println(is)
}

type Stack struct {
	stack []rune
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]rune, 0, 10),
	}
}

func (s *Stack) Push(a rune) {
	s.stack = append(s.stack, a)
}

func (s *Stack) Pop() rune {
	str := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return str
}

func (s *Stack) Top() rune {
	return s.stack[len(s.stack)-1]
}

func isValid(s string) bool {
	stack := NewStack()
	hash := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	// var z rune
	z := 0
	for _, b := range s {

		switch b {
		case '[', '(', '{':
			stack.Push(b)
		default:
			if len(stack.stack) > 0 {
				s := stack.Top()
				if v, ok := hash[b]; ok {
					if s == v {
						stack.Pop()

					} else {
						return false
					}
				}

			} else {
				z++
			}
		}

	}

	return len(stack.stack) == 0 && z == 0
}
