/*
 * Stack Algorithms Test
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

package stack

import "testing"

func Test_Stack_LinkedList(t *testing.T) {
	stack := NewStackLinkedList()

	if !stack.isEmpty() {
		t.Errorf("expected true")
	}

	stack.push("to")
	stack.push("be")
	stack.push("or")
	stack.push("not")
	stack.push("to")

	if stack.pop() != "to" {
		t.Errorf("expected 'to' ")
	}

	stack.push("be")

	if stack.pop() != "be" {
		t.Errorf("expected 'be' ")
	}

	if stack.pop() != "not" {
		t.Errorf("expected 'not' ")

	}

	if stack.pop() != "or" {
		t.Errorf("expected 'or' ")

	}
	stack.push("that")
	stack.push("is")

	if stack.pop() != "is" {
		t.Errorf("expected 'is' ")

	}

}
