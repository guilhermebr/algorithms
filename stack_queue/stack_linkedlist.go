/*
 * Stack Algorithms
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

package stack

type Node struct {
	item string
	next *Node
}

type StackLinkedList struct {
	first *Node
}

func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{nil}

}

func (s *StackLinkedList) isEmpty() bool {
	return s.first == nil
}

func (s *StackLinkedList) push(item string) {
	oldfirst := s.first
	s.first = &Node{}
	s.first.item = item
	s.first.next = oldfirst
}

func (s *StackLinkedList) pop() string {
	item := s.first.item
	s.first = s.first.next
	return item
}
