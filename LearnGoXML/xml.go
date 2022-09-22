package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	//"os"
)

func main() {
	type Contact struct {
		ID        int      `xml:"id,omitempty"`
		Name      string   `xml:"name,attr"`
		Age       int      `xml:"age"`
		Salary    float64  `xml:"salary"`
		IsMarried bool     `xml:"married"`
		Tags      []string `xml:"tags>tag"`
	}

	type Customer struct {
		ID          int     `xml:"id,omitempty"`
		Credit      float64 `xml:"credit"`
		ContactData Contact `xml:"contact"`
	}

	// XML with structs
	var c Contact = Contact{1, "Rafael", 50, 2022.9, true, []string{"Architect", "IT", "CTO", "Go"}}
	var cu Customer = Customer{ID: 1, Credit: 1000.0, ContactData: c}
	fmt.Println(cu)

	//b, err := xml.Marshal(c)
	//b, err := xml.Marshal(&c) // Marshal also supports references to struct
	b, err := xml.Marshal(&cu)
	if err != nil {
		fmt.Println("Error found: ", err)
		return
	}

	fmt.Println("XML cu: ", string(b))

	// XML with slices
	s := []Contact{c, {2, "Edu", 52, 1970.8, false, []string{"Operator", "CISCO"}}}

	bb, err := xml.MarshalIndent(&s, " ", "    ")
	if err != nil {
		fmt.Println("Error found: ", err)
		return
	}

	fmt.Println("XML cu: ", xml.Header, string(bb))

	// XML with maps *** UNSOPPORTED ***

	// Writin XML to a file
	f, err := os.Create("contacts.xml")
	PrintFatalError(err)
	defer f.Close()

	enc := xml.NewEncoder(f)
	enc.Indent(" ", "   ")
	enc.Encode(cu)

	PrintFatalError(err)

	// UnMarshal from a JSON format to struct
	sbyte := []byte(` <contact name="Rafael">
       <id>1</id>
       <age>50</age>
       <salary>2022.9</salary>
       <married>true</married>
       <tags>
          <tag>Architect</tag>
          <tag>IT</tag>
          <tag>CTO</tag>
          <tag>Go</tag>
       </tags>
    </contact>`)
	c2 := new(Contact)
	xml.Unmarshal(sbyte, &c2)
	fmt.Println("c2:", c2.ID, c2.Name, c2.Age, c2.Salary)

	// map is NO SUPPORTED for unmarshal

	// JSON from a file to memory
	/*
		file, err := os.Open("contacts.xml")
		PrintFatalError(err)
		defer file.Close()
		var c3 interface{}
		xml.NewDecoder(file).Decode(c3)
		fmt.Println("c3:", c3)
	*/
	var c3 interface{}
	err = decodeXML(c3, "contacts2.xml")
	PrintFatalError(err)
	fmt.Println("c3:", c3)

}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error has occur:", err)
	}
}

func decodeXML(v interface{}, filename string) error {
	fmt.Println("Decoding XML...")
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	return xml.NewDecoder(file).Decode(&v)
}
