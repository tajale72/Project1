package main

import "fmt"

func main() {

}

func Array() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}

type Node struct {
	data int
	next *Node
}

func LinkedList() {
	head := &Node{data: 1}
}
