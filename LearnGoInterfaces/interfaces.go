package main

import "fmt"

// Interface declaration
type Node interface {
	SetValue(v int)
	GetValue() int
}

// type SLLNode
type SLLNode struct {
	next  *SLLNode
	value int
}

// Implementing methods interface
func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

// Implementing methods interface
func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

// Constructor
func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

// type PowerNode
type PowerNode struct {
	next  *PowerNode
	value int
}

// Implementing methods interface
func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

// Implementing methods interface
func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

// Constructor
func NewPowerNode() *PowerNode {
	return new(PowerNode)
}

// Main method
func main() {
	var node Node
	node = NewSLLNode()
	node.SetValue(4)
	fmt.Println("Value of node is ", node.GetValue())

	node = NewPowerNode()
	node.SetValue(5)
	fmt.Println("Value of node is ", node.GetValue())
}
