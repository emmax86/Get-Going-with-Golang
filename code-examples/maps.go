package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
    m["ASU"] = Vertex{
        33.420534, -111.933983,
    }
	fmt.Printf("%v\n", m)
    delete(m, "ASU")
    fmt.Printf("%v\n", m)
}