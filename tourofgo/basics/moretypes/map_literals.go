package main

type Vertex struct {
	Lat, Long float64
}

type Cities struct {
	name1, name2 string
}

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

var s = map[int]Cities{
	1: {
		"Seattle", "Atlanta",
	},
	2: {
		"Vancouver", "NY",
	},
}

// func main() {
// 	fmt.Println(m["Bell Labs"])
// 	fmt.Println(m)
// 	fmt.Println(s)
// 	fmt.Println(s[2])
// }
