package examples

import (
	"fmt"

	"../ds"
)

type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

func Bfs(tree Node) {
	q := ds.Queue{}
	q.Enqueue(tree)
	for !q.IsEmpty() {
		n, ok := q.Dequeue().(Node)
		if ok {
			fmt.Println(n.data)
			if n.left != nil && n.right != nil {
				q.Enqueue(*n.left)
				q.Enqueue(*n.right)
			}
		}
	}
}
func Dfs(tree *Node) {
	if tree != nil {
		fmt.Println(tree.data)
		Dfs(tree.left)
		Dfs(tree.right)

	}
}

func GetTreeData() Node {
	return Node{left: &Node{left: &Node{nil, "C", nil}, data: "A", right: &Node{nil, "D", nil}}, data: "Root", right: &Node{left: &Node{nil, "E", nil}, data: "B", right: &Node{nil, "F", nil}}}
}
