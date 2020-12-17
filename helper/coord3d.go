package helper

import "fmt"

type Coord3D struct {
	X int
	Y int
	Z int
}

func (c Coord3D) GetNeighbours() map[Coord3D]bool {
	nums := make([][]int, 0)
	for _, n := range []int{c.X,c.Y,c.Z} {
		nums = append(nums, []int{n-1, n, n+1})
	}
	fmt.Println(nums)
	allm := make(map[Coord3D]bool, 0)
	for _, p := range nums {
		np := cartesianProduct(p)
		for {
			product := np()
			if len(product) == 0 {
				break
			}
			cc := Coord3D{product[0], product[1], product[2]}
			allm[cc] = true
		}
	}

	return allm
}

func cartesianProduct(a []int) func() []int {
	p := make([]int, len(a))
	x := make([]int, len(p))
	return func() []int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = a[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(a) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}