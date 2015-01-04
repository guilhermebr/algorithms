/*
 * Union Find Algorithms test
 * @gbrezende
 * guilhermebr@gmail.com
 * http://guilhermebr.com
 */

package unionfind

import "testing"
import "fmt"

func Test_Quick_Find(t *testing.T) {
	qf := NewQuickFind(10)

	qf.union(9, 7)
	qf.union(5, 1)
	qf.union(7, 0)
	qf.union(8, 7)
	qf.union(7, 6)
	qf.union(1, 9)

	fmt.Println(qf.elements)

	if qf.connected(4, 3) == true {
		t.Errorf("expected false")
	}

	if qf.connected(6, 5) == false {
		t.Errorf("expected true")
	}

}

func Test_Quick_Union(t *testing.T) {
	qu := NewQuickUnion(10)

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

	if qu.connected(9, 2) == false {
		t.Errorf("expected true")
	}

	if qu.connected(6, 5) == false {
		t.Errorf("expected true")
	}

}

func Test_Quick_Union_Weight(t *testing.T) {
	quw := NewQuickUnionWeight(10)

	quw.union(6, 9)
	quw.union(4, 3)
	quw.union(3, 0)
	quw.union(2, 8)
	quw.union(5, 4)
	quw.union(0, 1)
	quw.union(6, 2)
	quw.union(0, 9)
	quw.union(6, 7)

	fmt.Println(quw.elements)

	if quw.connected(9, 2) == false {
		t.Errorf("expected true")
	}

	if quw.connected(6, 5) == false {
		t.Errorf("expected true")
	}

}
