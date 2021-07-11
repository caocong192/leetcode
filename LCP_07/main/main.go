package main

import "fmt"

func numWays(n int, relation [][]int, k int) (ans int) {
	edges := make([][]int, n)
	for _, r := range relation {
		src, dst := r[0], r[1]
		edges[src] = append(edges[src], dst)
	}

	var dfs func(int, int)
	dfs = func(m int, step int) {
		if step == k {
			if m == n-1 {
				ans++
			}
			return
		}

		for _, next := range edges[m] {
			fmt.Println(step, next)
			dfs(next, step+1)
		}
	}
	dfs(0, 0)
	return
}

func main() {
	n := 5
	relation := [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}
	k := 3
	ans := numWays(n, relation, k)
	fmt.Println(ans)

}
