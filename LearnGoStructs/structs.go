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

	cm := CMember{"Elizabeth", 96, "Backingham Palace", "Queen", 5}

	var cmember CMember
	cmember.name = "Elizabeth"
	cmember.age = 96
	cmember.address = "Backingham Palace"
	cmember.rank = "Queen"
	cmember.clearance = 5

	cmp := &cmember
	cmp.clearance = 10

	var crew []CMember
	crew = append(crew, cm, cmember, CMember{"Rafael", 51, "Arroyomolinos", "Technologist", 5})

	for i, v := range crew {
		fmt.Println(i, v)
	}

}
