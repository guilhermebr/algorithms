/*
 * Stack Algorithms
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

func NewStackArray(capacity int) *StackArray {
	return &StackArray{make([]string, capacity), 0}
}

func (s *StackArray) isEmpty() bool {
	return s.N == 0
}

func (s *StackArray) push(item string) {
	s.items[s.N] = item
	s.N++
}

func (s *StackArray) pop() string {
	s.N--
	return s.items[s.N]
}

func (s *StackArray) debug() {
	fmt.Println(s.N)
	for n, item := range s.items {
		fmt.Printf("%d : %s \n", n, item)
	}
}
