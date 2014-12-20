
"""
Wighted Quick Union with Path Compression Algorithm to solve Union Find Problem (connectivity problem) 
Solution: O(lg N)
@gbrezende
guilhermebr@gmail.com
http://guilhermebr.com
"""

class QuickUnionUF(object):
	def __init__(self, n):
		self.elements = range(0, n)
		self.size = [1 for i in range(0, n)]

	def root(self, i):
		while i != self.elements[i]:
		#	self.elements[i] = self.elements[self.elements[i]]
			i = self.elements[i]

		return i

	def connected(self, p, q):
		return self.root(p) == self.root(q)

	def union(self, p, q):
		root_p = self.root(p)
		root_q = self.root(q)
 		if (root_p == root_q):
 			return

		if self.size[root_p] < self.size[root_q]:
			self.elements[root_p] = root_q
			self.size[root_q] += self.size[root_p]

		else:
			self.elements[root_q] = root_p
			self.size[root_p] += self.size[root_q]

	def debug(self):

		print ' ' + str(self.elements)
		#print 'Size: ' + str(self.size)


if __name__ == '__main__':
	qu = QuickUnionUF(10)

#    6-9 4-3 3-0 2-8 5-4 0-1 6-2 0-9 6-7 

	qu.union(6, 9)
	qu.debug()

	qu.union(4, 3)
	qu.debug()

	qu.union(3, 0)
	qu.debug()
	
	qu.union(2, 8)
	qu.debug()

	qu.union(5, 4)
	qu.debug()

	qu.union(0, 1)
	qu.debug()

	qu.union(6, 2)
	qu.debug()

	qu.union(0, 9)
	qu.debug()

	qu.union(6, 7)
	qu.debug()
	print(qu.connected(9, 2))
	print(qu.connected(6, 5))