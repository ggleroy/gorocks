package main

func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		result := first
		first, second = second, first+second
		return result
	}
}

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }
