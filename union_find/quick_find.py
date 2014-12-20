"""
Quick Find Algorithm to solve Union Find Problem (connectivity problem) 
@gbrezende
guilhermebr@gmail.com
http://guilhermebr.com
"""

class QuickFindUF(object):
	def __init__(self, n):
		self.elements = range(0, n)

	def compare(self, p, q):
		return self.elements[p] == self.elements[q]

	def union(self, p, q):
		pid = self.elements[p]
		qid = self.elements[q]

		for i in range(0, len(self.elements)):
			if pid == self.elements[i]:
				self.elements[i] = qid


if __name__ == '__main__':
	qf = QuickFindUF(10)

	qf.union(9, 7) 
	qf.union(5, 1)
	qf.union(7, 0)
	qf.union(8, 7)
	qf.union(7, 6)
	qf.union(1, 9)
	print qf.elements
	print(qf.compare(4, 3))
	print(qf.compare(6, 5))