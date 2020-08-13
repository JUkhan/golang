package main

import (
	"fmt"

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
}
