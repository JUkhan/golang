package examples

import (
	"fmt"

	"../ds"
)

type EncodeData struct {
	EncodedMsg string
	Table      map[rune]charTimes
}
type charTimes struct {
	char  rune
	times int
	code  string
}

func getCharTimes(item ds.Item) (res charTimes) {
	switch item.(type) {
	case charTimes:
		res = item.(charTimes)
		break

	case Node:
		res = item.(Node).data.(charTimes)
		break
	}
	return
}

func getMinHeap(msg string) (minHeap *ds.PriorityQueue) {
	dic := map[rune]int{}
	for _, v := range msg {
		dic[v]++
	}
	arr := []ds.Item{}
	for key, val := range dic {
		arr = append(arr, charTimes{char: key, times: val})
	}

	minHeap = ds.NewPriorityQueue(func(pItem, cItem ds.Item) bool {
		return getCharTimes(pItem).times > getCharTimes(cItem).times
	})
	minHeap.Heapify(arr)
	return
}
func treeTraversal(tree *Node) (dic map[rune]charTimes) {
	dic = make(map[rune]charTimes)
	var dfs func(*Node, string)
	dfs = func(node *Node, str string) {
		item := getCharTimes(node.data)
		if item.char != 0 {
			item.code = str
			dic[item.char] = item
			fmt.Printf("%c - %s\n", item.char, str)
		} else {
			if node.left != nil {
				dfs(node.left, str+"0")
			}
			if node.right != nil {
				dfs(node.right, str+"1")
			}
		}
	}
	dfs(tree, "")
	return
}
func HuffmanCoding(msg string) (res EncodeData) {

	minHeap := getMinHeap(msg)

	for !minHeap.IsEmpty() {
		item1 := minHeap.Dequeue()
		if minHeap.IsEmpty() {
			tree := item1.(Node)
			res.Table = treeTraversal(&tree)
			for _, v := range msg {
				res.EncodedMsg += res.Table[v].code
			}
		} else {
			item2 := minHeap.Dequeue()
			v1 := getCharTimes(item1)
			v2 := getCharTimes(item2)
			node := Node{data: charTimes{times: v1.times + v2.times}}

			switch item1.(type) {
			case charTimes:
				node.left = &Node{data: item1.(charTimes)}
				break

			case Node:
				l := item1.(Node)
				node.left = &l
				break
			}
			switch item2.(type) {
			case charTimes:
				node.right = &Node{data: item2.(charTimes)}
				break

			case Node:
				r := item2.(Node)
				node.right = &r
				break
			}
			minHeap.Enqueue(node)
		}
	}

	return
}
