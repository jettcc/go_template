// 并查集模板
package dsu

type Dsu struct {
	root, size []int
	count      int
}

func NewDsu(n int) *Dsu {
	root, size := make([]int, n), make([]int, n)
	return &Dsu{
		root:  make([]int, n),
		size:  make([]int, n),
		count: n,
	}
}

func (d *Dsu) Find(x int) int {
	if x != d.root[x] {
		d.root[x] = d.Find(d.root[x])
	}
	return d.root[x]
}

func (d *Dsu) Union(x, y int) {
	rx, ry := d.Find(x), d.Find(y)
	if rx == ry {
		return
	}
	d.root[rx] = ry
	d.count--
}

func (d *Dsu) Connect(x, y int) bool {
	return d.Find(x) == d.Find(y)
}

func (d *Dsu) Count() int {
	return d.count
}
