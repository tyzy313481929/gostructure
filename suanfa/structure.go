package main

import (
	"fmt"
	"structure"
)

func main() {
	head := LinkedList.NewLinkedList()
	for i := 1; i < 5; i++ {
		node := LinkedList.NewINode(i, nil)
		head.Append(node)
	}
	//	head.PrintList()
	findNode, err := head.Find(3)
	if err != true {
		fmt.Println(err)
	} else {
		fmt.Println(findNode.X)
	}
	head.Remove(1)
	head.PrintList()
}
