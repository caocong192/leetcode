package alertNames

import "sort"

func alertNames(keyName []string, keyTime []string) (want []string) {
	// 创建 员工-时间 map表
	timeMap := make(map[string][]int)

	for i, name := range keyName {
		// 时间转化为分钟 int类型
		t := keyTime[i]
		h := int(t[0]-'0')*10 + int(t[1]-'0')
		m := int(t[3]-'0')*10 + int(t[4]-'0')
		timeMap[name] = append(timeMap[name], h*60+m)
	}

	for p, t := range timeMap {
		// 如果打卡记录小于3, 跳过
		if len(t) < 3 {
			continue
		}

		sort.Ints(t)
		for i := 1; i < len(t)-1; i++ {
			if t[i+1]-t[i-1] <= 60 && t[i+1]-t[i-1] >= 0 {
				want = append(want, p)
				break
			}
		}
	}
	sort.Strings(want)
	return
}
