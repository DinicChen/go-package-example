package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var nilmap map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	delete(m, "key")
	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967}]

	delete(m, "Bell Labs")
	fmt.Println(m) // map[]

	delete(nilmap, "key")
}
