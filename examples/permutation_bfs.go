package examples

import (
	"../ds"
)

type state struct {
	values []rune
	lavel  int
}

//permutation using BFS(branch and bound) FIFI BB
func PermutationBfs(str string) (res []string) {
	arr := []rune(str)
	q := ds.Queue{}
	n := len(arr)
	for i := 0; i < n; i++ {
		q.Enqueue(state{values: []rune{arr[i]}, lavel: 1})
	}
	foundInValues := func(arr []rune, val rune, n int) bool {
		for j := 0; j < n; j++ {
			if arr[j] == val {
				return true
			}
		}
		return false
	}
	for !q.IsEmpty() {
		d, _ := q.Dequeue().(state)
		lenValues := len(d.values)
		for i := 0; i < n; i++ {
			if !foundInValues(d.values, arr[i], lenValues) {
				q.Enqueue(state{values: append(d.values, arr[i]), lavel: d.lavel + 1})
			}

		}
		if d.lavel == n {
			res = append(res, string(d.values))
		}
	}
	return
}
