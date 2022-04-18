// Question 4: Create a Stack to Reverse a String
package main

import "fmt"

type Stack struct {
	data []string
}

func (s *Stack) Push(value string) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() string {
	val := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]
	return val
}

func (s Stack) Read() string {
	return s.data[len(s.data) - 1]
}

func reverseString(str string) string {
	stack := new(Stack)

	for i := range str {
		s := string(str[i])
		stack.Push(s)
	}

	result := ""
	for len(stack.data) > 0 {
		result += stack.Pop()
	}

	return result
}

func main() {
	fmt.Println(reverseString("abcde"))
}