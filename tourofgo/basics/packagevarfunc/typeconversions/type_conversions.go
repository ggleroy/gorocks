package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	var jj int = 42
	var gg float64 = float64(jj)
	var uu uint = uint(gg)
	fmt.Println(jj, gg, uu)
	//Or simply like this
	j := 42
	g := float64(j)
	u := uint(g)
	fmt.Println(j, g, u)
}
