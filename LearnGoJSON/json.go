package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	type Contact struct {
		ID        int      `json:"id,omitempty"`
		Name      string   `json:"name"`
		Age       int      `json:"age"`
		Salary    float64  `json:"salary"`
		IsMarried bool     `json:"married"`
		Tags      []string `json:"tags"`
	}

	type Customer struct {
		ID          int     `json:"id,omitempty"`
		Credit      float64 `json:"credit"`
		ContactData Contact `json:"contact"`
	}

	// JSON with structs
	var c Contact = Contact{1, "Rafael", 50, 2022.9, true, []string{"Architect", "IT", "CTO", "Go"}}
	var cu Customer = Customer{ID: 1, Credit: 1000.0, ContactData: c}
	fmt.Println(cu)

	//b, err := json.Marshal(c)
	//b, err := json.Marshal(&c) // Marshal also supports references to struct
	b, err := json.Marshal(&cu)
	if err != nil {
		fmt.Println("Error found: ", err)
		return
	}

	fmt.Println("JSON cu: ", string(b))

	// JSON with slices
	s := []Contact{c, {2, "Edu", 52, 1970.8, false, []string{"Operator", "CISCO"}}}

	bb, err := json.Marshal(&s)
	if err != nil {
		fmt.Println("Error found: ", err)
		return
	}

	fmt.Println("JSON cu: ", string(bb))

	// JSON with maps
	m := map[string]int{"item1": 1, "item2": 2}
	bm, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error found: ", err)
		return
	}
	fmt.Println("JSON m: ", string(bm))

	// Writin JSON to a file
	f, err := os.Create("contacts.json")
	PrintFatalError(err)
	defer f.Close()

	err = json.NewEncoder(f).Encode(&cu)
	PrintFatalError(err)

	// UnMarshal from a JSON format to struct
	sbyte := []byte(`{"id":1,"name":"Rafael","age":50,"salary":2022.9,"married":true,"tags":["Architect","IT","CTO","Go"]}`)
	c2 := new(Contact)
	json.Unmarshal(sbyte, c2)
	fmt.Println("c2:", c2.ID, c2.Name, c2.Age, c2.Salary)

	m2 := make(map[int]string)
	sbyte = []byte(`{"1":"item1","2":"item2","3":"item3","4":"item4"}`)
	json.Unmarshal(sbyte, &m2)
	fmt.Println("m2:", m2)

	sbyte = []byte(`[{"id":1,"name":"Rafael","age":50,"salary":2022.9,"married":true,"tags":["Architect","IT","CTO","Go"]},{"id":2, "name":"Edu", "age":52, "salary":1970.8, "married":false, "tags":["Operator", "CISCO"]}]`)
	s2 := []Contact{}
	json.Unmarshal(sbyte, &s2)
	fmt.Println("s2:", s2)

	// JSON from a file to memory
	file, err := os.Open("contacts.json")
	PrintFatalError(err)
	defer file.Close()
	var c3 interface{}
	json.NewDecoder(file).Decode(&c3)
	fmt.Println("c3:", c3)
}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error has occur:", err)
	}
}
