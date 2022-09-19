package main

import "fmt"

type CMember struct {
	name      string
	age       int
	address   string
	rank      string
	clearance int
}

func main() {

	// Declaring a map
	var m map[string]CMember
	m = make(map[string]CMember)

	// Adding a regster to the map using a Key
	cm := CMember{"Elizabeth", 96, "Backingham Palace", "Queen", 5}
	m["Elizabeth"] = cm

	// Adding a regster to the map using a Key
	m["Rafael"] = CMember{"Rafael", 51, "Arroyomolinos", "Technologist", 5}

	// Accessing to each register using the key
	fmt.Println(m["Elizabeth"])
	fmt.Println(m["Rafael"].rank)

	// Accessing to every register in the map
	for i, v := range m {
		fmt.Println("[", i, "]", v)
	}

	// Assigning several registers in a map
	m2 := map[string]CMember{
		"Nerea": CMember{name: "Nerea", age: 20, address: "Paris", rank: "Interpreter", clearance: 3},
		"Eddy":  CMember{name: "Eddy", age: 52, address: "Mostoles", rank: "IT Operator", clearance: 4},
	}

	// Accessing to every register in the map
	for i, v := range m2 {
		fmt.Println("[", i, "]", v)
	}

	// Checking if a key exists
	v, ok := m2["Nerea"] // v=register, ok=true
	fmt.Println(v, ok)

	// Deleting a register in the map using its key
	delete(m2, "Eddy")
	v, ok = m2["Edit"] // v=0, ok=false
	fmt.Println(v, ok)

}
