package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	fmt.Println("Begin sum")
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("Writing to channel")
	c <- sum // send sum to c
	fmt.Println("end sum")
}

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int, 2)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	fmt.Println("lendo x")
// 	x := <-c
// 	fmt.Println("lido x")
// 	time.Sleep(time.Millisecond * 5000)
// 	fmt.Println("lendo y")
// 	y := <-c // receive from c
// 	fmt.Println("lido y")

// 	fmt.Println(x, y, x+y)
// }
