package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "Mundo"
	fmt.Println("Olá,", World)
	fmt.Println("Happy", Pi, "Day!")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
