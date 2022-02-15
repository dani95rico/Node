// EXAMPLE OF MULTIPLES NODES IN GO LENGUAGE BY DANI95RICO

package main

import "fmt"

// Node Interface
type Node interface {
	SetValue(v int)
	GetValue() int
}

// Type SLLNode
type SLLNode struct {
	next  *SLLNode
	value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

// Type PowerNode
type PowerNode struct {
	next  *PowerNode
	value int
}

func (pNode *PowerNode) SetValue(v int) {
	pNode.value = v * 10
}

func (pNode *PowerNode) GetValue() int {
	return pNode.value
}

func NewPowerNode() *PowerNode {
	return new(PowerNode)
}

type SingleLinkedList struct {
	cabeza *SLLNode
	cola   *SLLNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

func (list *SingleLinkedList) Add(v int) {
	newNode := &SLLNode{value: v}
	if list.cabeza == nil {
		list.cabeza = newNode
	} else if list.cola == list.cabeza {
		list.cabeza.next = newNode
	} else if list.cola != nil {
		list.cola.next = newNode
	}
	list.cola = newNode
}

func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.cabeza; n != nil; n = n.next {
		s += fmt.Sprintf("{%d}", n.GetValue())
	}
	return s
}

func main() {
	var node Node
	node = NewSLLNode()
	node.SetValue(5)
	fmt.Println("Value of the node: ", node.GetValue())

	node = NewPowerNode()
	node.SetValue(5)
	fmt.Println("Value of the node: ", node.GetValue())

	if n, ok := node.(*PowerNode); ok {
		fmt.Println("This is a PowerNode with a value of: ", n.value)
	} else {
		fmt.Println("This is a SLLNode with a value of: ", n.value)
	}

	list := NewSingleLinkedList()
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	fmt.Println("Value of the node list: ", list)
}
