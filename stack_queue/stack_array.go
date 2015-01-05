/*
 * Stack Array Algorithms
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

package stack

import "fmt"

type StackArray struct {
	items []string
	N     int
}

func NewStackArray() *StackArray {
	return &StackArray{make([]string, 1), 0}
}

func (s *StackArray) isEmpty() bool {
	return s.N == 0
}

func (s *StackArray) push(item string) {
	if s.N == len(s.items) {
		s.resize(2 * len(s.items))
	}

	s.items[s.N] = item
	s.N++
}

func (s *StackArray) pop() string {
	s.N--
	item := s.items[s.N]
	s.items[s.N] = ""

	if s.N > 0 && s.N == (len(s.items)/4) {
		s.resize(len(s.items) / 2)
	}
	return item
}

func (s *StackArray) resize(capacity int) {
	fmt.Printf("Resizing to: %d\n", capacity)
	copy := make([]string, capacity)
	for i := 0; i < s.N; i++ {
		copy[i] = s.items[i]
	}
	s.items = copy
	s.debug()
}

func (s *StackArray) debug() {
	fmt.Printf("N: %d\n", s.N)
	fmt.Printf("Array size: %d\n", len(s.items))
	for n, item := range s.items {
		fmt.Printf("%d : %s \n", n, item)
	}
	fmt.Println()
}
