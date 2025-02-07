package ds

type Stack struct {
	Items []int
}

type Queue struct {
	Items []int
}

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}
