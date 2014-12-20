/*
 * Quick Union Algorithm to solve Union Find Problem (connectivity problem)
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

package main

import "fmt"

type QuickUnionUF struct {
	elements []int
}

func NewQuickUnionUF(N int) *QuickUnionUF {
	qu := &QuickUnionUF{}

	qu.elements = make([]int, N)

	for i := 0; i < N; i++ {
		qu.elements[i] = i
	}

	return qu
}

func (self *QuickUnionUF) root(i int) int {
	for i != self.elements[i] {
		i = self.elements[i]

	}
	return i
}

func (self *QuickUnionUF) connected(p, q int) bool {
	return self.root(p) == self.root(q)
}

func (self *QuickUnionUF) union(p, q int) {
	self.elements[self.root(p)] = self.root(q)
}

func main() {

	qu := NewQuickUnionUF(10)

	//  8-7 5-6 6-0 0-8 2-1 3-4 3-1 4-7 4-9
	qu.union(8, 7)
	qu.union(5, 6)
	qu.union(6, 0)
	qu.union(0, 8)
	qu.union(2, 1)
	qu.union(3, 4)
	qu.union(3, 1)
	qu.union(4, 7)
	qu.union(4, 9)

	fmt.Println(qu.elements)
	//	fmt.Println(qu.connected(9, 2))
	//	fmt.Println(qu.connected(6, 5))
}
