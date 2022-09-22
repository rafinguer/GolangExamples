package main

import (
	"fmt"
	"reflect"
)

type contact struct {
	Name      string  `alias:"fullName" desc:"Full name of the contact"`
	Age       int     `alias:"Age" desc:"Age (years old) of the contact"`
	Salary    float64 `alias:"Salary" desc:"Salary (in euros) of the contact"`
	IsMarried bool    `alias:"Married" desc:"Is the contact married? (true/false)"`
}

func main() {
	var c contact
	c.Name = "Rafael"
	c.Age = 50
	c.Salary = 2022.9
	c.IsMarried = true

	v := reflect.ValueOf(c)
	t := reflect.TypeOf(c)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // datatype: StructField
		value := v.Field(i)
		fmt.Println("---")
		fmt.Printf("Field name: %s - Field type: %s - Field value: %v\n", field.Name, field.Type, value.Interface())
		fmt.Println("   alias:", field.Tag.Get("alias"), ", description: ", field.Tag.Get("desc"))
	}
}
