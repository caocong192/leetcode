package hardestWorker

func hardestWorker(n int, logs [][]int) int {
	ans, max := logs[0][0], logs[0][1]

	start := 0
	for _, log := range logs {
		if log[1]-start > max || log[1]-start == max && log[0] < ans {
			max = log[1] - start
			ans = log[0]
		}
		start = log[1]
	}
	return ans
}
