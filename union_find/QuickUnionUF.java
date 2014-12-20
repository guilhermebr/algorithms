/*
 * Quick Union Algorithm to solve Union Find Problem (connectivity problem)
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

public class QuickUnionUF {

	private int[] elements;

	public QuickUnionUF(int N) {

		elements = new int[N];

		for (int i = 0; i < N; i++) {
			elements[i] = i;
		}

	}

	private int root(int i) {
		while (i != elements[i] ) {
			i = elements[i];
		}
		return i;
	}

	public boolean connected(int p, int q) {
		return root(p) == root(q);
	}

	public void union(int p, int q) {
		elements[root(p)] =  root(q);
	}

	public static void main(String []args) {
		boolean x;
	    QuickUnionUF qu = new QuickUnionUF(10);
		qu.union(3, 4);
		qu.union(4, 9);
		qu.union(2, 9);
		qu.union(5, 6);
		qu.union(3, 5);		
		x = qu.connected(4, 3);
		System.out.println(x);
		x = qu.connected(6, 4);
		System.out.println(x);

	}
}