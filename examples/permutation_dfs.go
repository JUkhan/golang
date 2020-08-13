package examples

//permutation using backtracking
func PermutationDfs(str string) (res []string) {
	arr := []rune(str)
	var helper func(i, n int)
	helper = func(i, n int) {
		if i == n {
			res = append(res, string(arr))
		} else {
			for j := i; j < n; j++ {
				arr[i], arr[j] = arr[j], arr[i]
				helper(i+1, n)
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	helper(0, len(arr))
	return
}
