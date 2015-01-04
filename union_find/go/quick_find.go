/*
 * Quick Find Algorithm to solve Union Find Problem (connectivity problem)
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

package unionfind

type QuickFind struct {
	elements []int
}

func NewQuickFind(N int) *QuickFind {
	qf := &QuickFind{}

	qf.elements = make([]int, N)

	for i := 0; i < N; i++ {
		qf.elements[i] = i
	}

	return qf

}

func (self *QuickFind) connected(p, q int) bool {
	return self.elements[p] == self.elements[q]
}

func (self *QuickFind) union(p, q int) {
	pid := self.elements[p]
	qid := self.elements[q]

	for i := 0; i <= 9; i++ {
		if self.elements[i] == pid {
			self.elements[i] = qid
		}
	}

}
