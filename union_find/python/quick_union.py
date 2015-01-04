
"""
Quick Union Algorithm to solve Union Find Problem (connectivity problem) 
@gbrezende
guilhermebr@gmail.com
http://guilhermebr.com
"""

class QuickUnionUF(object):
	def __init__(self, n):
		self.elements = range(0, n)

	def root(self, i):
		while i != self.elements[i]:
			i = self.elements[i]

		return i

	def connected(self, p, q):
		return self.root(p) == self.root(q)

	def union(self, p, q):
		self.elements[self.root(p)] = self.root(q)


if __name__ == '__main__':
	qu = QuickUnionUF(10)

	qu.union(3, 4)
	print qu.elements
	qu.union(4, 9)
	print qu.elements
	qu.union(2, 9)
	print qu.elements
	qu.union(5, 6)
	print qu.elements
	qu.union(3, 5)
	print qu.elements
	print(qu.connected(9, 2))
	print(qu.connected(6, 5))