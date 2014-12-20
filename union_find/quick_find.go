/*
 * Quick Find Algorithm to solve Union Find Problem (connectivity problem)
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

package main

import "fmt"

type QuickFindUF struct {
	elements []int
}

func NewQuickFindUF(N int) *QuickFindUF {
	qf := &QuickFindUF{}

	qf.elements = make([]int, N)

	for i := 0; i < N; i++ {
		qf.elements[i] = i
	}

	return qf

}

func (self *QuickFindUF) compare(p, q int) bool {
	return self.elements[p] == self.elements[q]
}

func (self *QuickFindUF) union(p, q int) {
	pid := self.elements[p]
	qid := self.elements[q]

	for i := 0; i <= 9; i++ {
		if self.elements[i] == pid {
			self.elements[i] = qid
		}
	}

}

func main() {

	qf := NewQuickFindUF(10)

	qf.union(9, 7)
	qf.union(5, 1)
	qf.union(7, 0)
	qf.union(8, 7)
	qf.union(7, 6)
	qf.union(1, 9)

	fmt.Println(qf.elements)
	//	fmt.Println(qf.compare(4, 3))
	///	fmt.Println(qf.compare(6, 4))

}
