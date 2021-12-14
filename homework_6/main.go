package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type listNode struct {
	value int
	next  *listNode
}

func main() {
	var num int

	if _, err := fmt.Scanln(&num); errors.Is(err, io.EOF) {
		os.Exit(1)
	}

	var firstNode *listNode
	firstNode = &listNode{value: num}
	prevNode := firstNode

	// Filling firstNode (listNode)
	for {
		if _, err := fmt.Scanln(&num); errors.Is(err, io.EOF) {
			fmt.Println("Input completed")
			break
		}
		nextNode := &listNode{value: num}
		prevNode.next = nextNode
		prevNode = nextNode
	}

	// Revers firstNode
	var reversListNode *listNode
	prevRevNode := firstNode
	currentNode := firstNode.next
	firstNode.next = nil
	for {
		if currentNode == nil {
			reversListNode = prevRevNode
			break
		}
		nextNode := currentNode.next
		currentNode.next = prevRevNode
		prevRevNode = currentNode
		currentNode = nextNode
	}

	// Output on display ReversListNode
	for reversListNode != nil {
		fmt.Println(reversListNode.value)
		reversListNode = reversListNode.next
	}
}
