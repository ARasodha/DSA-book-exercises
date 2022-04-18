/*
Question 1:
Add a method to the classic LinkedList class that prints all of the elements of the list
Question 2: On chapter14-pt2.go due to using DoublyLinkedList
Question 3: Add a method to the class LinkedList class that returns the last element from the list.
Assume you don't know how many elements are in the list.
Question 4: Here's a tricky one. Add a method to the classic LinkedList class that reverses the list.
That is, if the original list is A -> B -> C, all of the list's links should change so that 
C -> B -> A
Question 5: Here's a brilliant little linked list puzzle for you. Let's say you have access to a
node from somewhere in the middle of a classic linked list, but not the linked list itself. That is,
you have a variable that points to an instance of Node, but you don't have access to the LinkedList
instance. In this situation, if you follow this node's link, you can find all the items from this
middle node until the end, but you have no way to find the nodes that precede this node in the list.
Write code that will effectively delete this node from the list. The entire remaining list should
remain complete, with only this node removed.
*/

package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}

func (l *LinkedList) insertNode(index int, value string) {
	newNode := &Node{data: value}

	if index == 0 {
		newNode.next = l.head
		l.head = newNode
	} else {
		currentNode := l.head
		currentIndex := 0
	
		for currentIndex < (index - 1) {
			currentNode = currentNode.next
			currentIndex += 1
		}
	
		newNode.next = currentNode.next
		currentNode.next = newNode
	}
	
	l.length++
}

func (l LinkedList) printAllNodes() {
	currentNode := l.head
	for i :=0; i<l.length; i++ {
		fmt.Println(currentNode)
		currentNode = currentNode.next
	}
}

func (l LinkedList) lastNode() *Node {
	currentNode := l.head
	for i := 0; i<l.length - 1; i++ {
		currentNode = currentNode.next
	}

	return currentNode
}

func (l *LinkedList) reverseList() {
	previousNode := new(Node)
	currentNode := l.head
	nextNode := currentNode

	for i:=0; i<l.length; i++ {
		nextNode = currentNode.next
		currentNode.next = previousNode

		previousNode = currentNode
		currentNode = nextNode
	}

	l.head = previousNode
}

func (l *LinkedList) deleteMiddleNode(n *node) {
	node.data = node.next.data
	node.next = node.next.next
}

func main() {
	myList := LinkedList{}
	myList.insertNode(myList.length, "Once")
	myList.insertNode(myList.length, "Upon")
	myList.insertNode(myList.length, "A")
	myList.insertNode(myList.length, "Time")
	// fmt.Println(myList)
	// myList.printAllNodes()
	fmt.Println("last node:", myList.lastNode())
	myList.reverseList()
	myList.printAllNodes()
}