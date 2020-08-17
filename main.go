package main

import (
	"fmt"

	"./ds"
	"./examples"
	//"github.com/cheekybits/genny/generic"
)

//type item generic.Type

//its a function take int as param and return string
type convert func(int) string

func test(data int, converter convert) string {

	return converter(data)
}
func value(d int) string {
	return fmt.Sprintf("value : %d", d)
}
func bar(n interface{}) int {
	return n.(int)
}
func main() {

	//http.HandleFunc("/", handle_func)
	//http.ListenAndServe(":9999", nil)
	fmt.Println(test(55, func(d int) string {
		return fmt.Sprintf("values : %d", d)
	}))

	tree := examples.GetTreeData()

	fmt.Println("-----------BFS----------")
	examples.Bfs(tree)

	fmt.Println("-----------DFS----------")
	examples.Dfs(&tree)
	s := bar(34)
	fmt.Println(s + 3)
	fmt.Println("-------------permutation using backtracking-----------------")
	fmt.Println(examples.PermutationDfs("ABC"))
	fmt.Println("-------------permutation using FIFO-BB(BFS)-----------------")
	fmt.Println(examples.PermutationBfs("ABC"))
	fmt.Println("-------------Knapsack using greedy method-----------------")
	kdata := []examples.KData{
		{Product: "1", Price: 10, Weight: 2},
		{Product: "2", Price: 5, Weight: 3},
		{Product: "3", Price: 15, Weight: 5},
		{Product: "4", Price: 7, Weight: 7},
		{Product: "5", Price: 6, Weight: 1},
		{Product: "6", Price: 18, Weight: 4},
		{Product: "7", Price: 3, Weight: 1},
	}
	fmt.Println(examples.Knapsack(kdata, 15))

	fmt.Println("-------------Activity Selection-----------------")
	adata := []examples.ActivetyData{
		{Activity: "A1", Start: 12, Finish: 25},
		{Activity: "A2", Start: 10, Finish: 20},
		{Activity: "A3", Start: 30, Finish: 30},
	}
	fmt.Println(examples.ActivitySelection(adata))

	fmt.Println("-------------PriorityQueue-----------------")
	pq := ds.NewPriorityQueue(func(pItem, cItem ds.Item) bool {
		p := pItem.(int)
		c := cItem.(int)
		return p < c
	})

	/*pq.Enqueue(15)
	pq.Enqueue(10)
	pq.Enqueue(30)
	pq.Enqueue(20)
	pq.Enqueue(50)
	pq.Enqueue(8)
	pq.Enqueue(16)
	pq.Enqueue(2)*/
	pq.Heapify([]ds.Item{15, 10, 30, 20, 50, 8, 16, 2})

	for !pq.IsEmpty() {
		fmt.Println(pq.Dequeue())
	}
	fmt.Println("-------------HuffmanCoding-----------------")
	fmt.Println("Original msg: EEAAADDDDBBBBBCCCCCC")
	fmt.Println("Encoded msg:", examples.HuffmanCoding("EEAAADDDDBBBBBCCCCCC").EncodedMsg)
}
