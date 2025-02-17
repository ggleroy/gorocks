package main

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// // removing the * changes the behavior of the method
// // from a pointer receiver (more common) to a value receiver
// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// 	v.Scale(10)
// 	fmt.Println(v.Abs())
// }
