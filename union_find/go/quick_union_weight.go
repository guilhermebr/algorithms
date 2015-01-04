/*
 * Wighted Quick Union with Path Compression Algorithm to solve Union Find Problem (connectivity problem)
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

package unionfind

type QuickUnionWeight struct {
	elements []int
	size     []int
}

func NewQuickUnionWeight(N int) *QuickUnionWeight {
	quw := &QuickUnionWeight{}

	quw.elements = make([]int, N)
	quw.size = make([]int, N)

	for i := 0; i < N; i++ {
		quw.elements[i] = i
		quw.size[i] = 1
	}

	return quw
}

func (self *QuickUnionWeight) root(i int) int {
	for i != self.elements[i] {
		i = self.elements[i]
	}

	return i
}

func (self *QuickUnionWeight) connected(p, q int) bool {
	return self.root(p) == self.root(q)
}

func (self *QuickUnionWeight) union(p, q int) {
	root_p := self.root(p)
	root_q := self.root(q)

	if root_p == root_q {
		return
	}

	if self.size[root_p] < self.size[root_q] {
		self.elements[root_p] = root_q
		self.size[root_q] = self.size[root_p]
	} else {
		self.elements[root_q] = root_p
		self.size[root_p] = self.size[root_q]
	}

}
