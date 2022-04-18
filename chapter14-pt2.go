/*
Question 2: Add a method to the DoublyLinkedList class that prints all the elements of the list
in reverse order
*/

package main

import "fmt"

type Node struct {
	data string
	next *Node
	previous *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	length int
}

func (d *DoublyLinkedList) insertAtEnd(value string) {
	newNode := &Node{data: value}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.previous = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}
	d.length++
}

func (d * DoublyLinkedList) printNodesReverse() {
	currentNode := d.tail
	for i:=0; i<d.length; i++ {
		fmt.Println(currentNode)
		currentNode = currentNode.previous
	}
}

func main() {
	dLL := DoublyLinkedList{}
	dLL.insertAtEnd("Once")
	dLL.insertAtEnd("Upon")
	dLL.insertAtEnd("A")
	dLL.insertAtEnd("Time")
	fmt.Println(dLL)
	dLL.printNodesReverse()
}