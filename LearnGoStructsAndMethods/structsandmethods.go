package main

import "fmt"

// type SLLNode
type SLLNode struct {
	next  *SLLNode
	value int
}

// With (sNode *SLLNode) we associate the method to the struct
func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

// With (sNode *SLLNode) we associate the method to the struct
func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

// Constructor
// Returns an instance of the struct, including methods
// Exatly, returns a reference to the memory address of struct
func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

// Level struct
type level int

// Method endorsed to level struct
func (lv *level) raiseShieldLevel(i int) {
	if *lv == 0 {
		*lv = 1
	}

	*lv = (*lv) * level(i) // level(i) casting i to level struct
}

// Main method
func main() {

	// Testing SSLNode struct and methods
	node := NewSLLNode()
	node.SetValue(12)
	fmt.Println("Value of node is ", node.GetValue())

	// Testing level struct and method
	sl := new(level)
	sl.raiseShieldLevel(4)
	sl.raiseShieldLevel(5)
	fmt.Println(*sl)

	var sl2 level = 5
	sl2.raiseShieldLevel(10)
	fmt.Println(sl2)
}
