package main

import (
	"fmt"
	"sync"
)

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Reader Leroy %d\n", id)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go reader(1, &wg)
	go reader(2, &wg)

	wg.Wait()
}
