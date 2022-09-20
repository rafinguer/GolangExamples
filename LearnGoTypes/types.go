package main

import (
	"fmt"
	"reflect"
	"strings"
)

// magicStore struct with methods
type magicStore struct {
	value interface{} // any type
	name  string
}

func (ms *magicStore) SetValue(v interface{}) {
	ms.value = v
}

func (ms *magicStore) GetValue() interface{} {
	return ms.value
}

/*
func (ms *magicStore) NewMagicStore() *magicStore {  // Constructor

		return new(magicStore)
	}
*/
func NewMagicStore(nm string) *magicStore { // Constructor with arguments
	return &magicStore{name: nm}
}

func main() {

	// Native types
	printType("text")
	printType(32)
	printType(23.56)
	printType(false)

	// Complex type: magicStore struct
	istore := NewMagicStore("Integer store")
	istore.SetValue(25)
	if v, ok := istore.GetValue().(int); ok {
		v *= 4
		fmt.Println("MagicStore value of int type: ", v)
	}

	sstore := NewMagicStore("String store")
	sstore.SetValue("Demo text")
	if v, ok := sstore.GetValue().(string); ok {
		v += "***"
		fmt.Println("MagicStore value of string type: ", v)
	}

	var stype string = reflect.TypeOf(sstore).String()

	if strings.Contains(stype, "magicStore") {
		fmt.Println("sstore is magicStore type: ", sstore)
	}
}

// This function prints the type of a given value
func printType(i interface{}) {
	switch i := i.(type) {
	case string:
		fmt.Println("Data is string: ", i)
	case int:
		fmt.Println("Data is int: ", i)
	case float64:
		fmt.Println("Data is float64: ", i)
	case bool:
		fmt.Println("Data is bool: ", i)
	}
}
