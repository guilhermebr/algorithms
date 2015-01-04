/*
 * Quick Union Algorithm to solve Union Find Problem (connectivity problem)
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

package unionfind

type QuickUnion struct {
	elements []int
}

func NewQuickUnion(N int) *QuickUnion {
	qu := &QuickUnion{}

	qu.elements = make([]int, N)

	for i := 0; i < N; i++ {
		qu.elements[i] = i
	}

	return qu
}

func (self *QuickUnion) root(i int) int {
	for i != self.elements[i] {
		i = self.elements[i]

	}
	return i
}

func (self *QuickUnion) connected(p, q int) bool {
	return self.root(p) == self.root(q)
}

func (self *QuickUnion) union(p, q int) {
	self.elements[self.root(p)] = self.root(q)
}
