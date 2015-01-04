/*
 * Quick Find Algorithm to solve Union Find Problem (connectivity problem)
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */


public class QuickFindUF {
	
	private int[] elements;
 	
 	public QuickFindUF(int N) {
 		elements = new int[N];
 		for (int i = 0; i < N; i++)
 			elements[i] = i;
 	}
 
	public boolean connected(int p, int q) { 
	 	return elements[p] == elements[q]; 
	 }

	public void union(int p, int q) {
		int pid = elements[p];
	 	int qid = elements[q];
	 	for (int i = 0; i < elements.length; i++)
	 		if (elements[i] == pid) elements[i] = qid;
	}

	public static void main(String []args) {
		boolean x;
	    QuickFindUF qf = new QuickFindUF(10);
		qf.union(4, 3);
		x = qf.connected(4, 3);
		System.out.println(x);
		x = qf.connected(6, 4);
		System.out.println(x);

	}

}