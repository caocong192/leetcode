package findMinimumTime

import "sort"

func findMinimumTime(tasks [][]int) (ans int) {
	// 贪心
	// 按照任务结束时间从短到大排序
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})

	n := len(tasks)
	run := make([]int, tasks[n-1][1]+1)

	for i := 0; i < n; i++ {
		start, end, duration := tasks[i][0], tasks[i][1], tasks[i][2]
		for j := start; j <= end; j++ {
			duration -= run[j]
		}

		if duration > 0 {
			for j := end; j >= start; j-- {
				if run[j] == 0 {
					run[j] = 1
					duration--
				}
				if duration == 0 {
					break
				}
			}
		}
	}

	for i := 0; i < len(run); i++ {
		if run[i] == 1 {
			ans++
		}
	}

	return
}
