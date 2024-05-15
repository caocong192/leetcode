package garbageCollection

func garbageCollection(garbage []string, travel []int) (ans int) {
	// 用于标识是否有对应的垃圾
	gFlag, pFlag, mFlag := false, false, false
	// 倒序遍历
	for i := len(garbage) - 1; i >= 0; i-- {
		for j := 0; j < len(garbage[i]); j++ {
			if garbage[i][j] == 'G' {
				gFlag = true
				garbage[i] = garbage[i][:j] + " " + garbage[i][j+1:]
				ans++

			}
			if garbage[i][j] == 'P' {
				pFlag = true
				garbage[i] = garbage[i][:j] + " " + garbage[i][j+1:]
				ans++

			}
			if garbage[i][j] == 'M' {
				mFlag = true
				garbage[i] = garbage[i][:j] + " " + garbage[i][j+1:]
				ans++
			}
		}

		// 如果开始有垃圾了, 且i不为0, 则添加对应的距离消耗
		if gFlag == true && i != 0 {
			ans += travel[i-1]
		}

		if pFlag == true && i != 0 {
			ans += travel[i-1]
		}

		if mFlag == true && i != 0 {
			ans += travel[i-1]
		}
	}
	return
}
