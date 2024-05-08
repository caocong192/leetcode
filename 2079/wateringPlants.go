package wateringPlants

func wateringPlants(plants []int, capacity int) (ans int) {
	// bucket 水桶初始值为 capacity
	bucket := capacity
	for i, v := range plants {
		if bucket >= v {
			bucket -= v
			ans++
		} else {
			// 如果水量不够, 需要走 2*i + 1步  取完水再来浇
			ans += 2*i + 1
			bucket = capacity - v
		}
	}
	return
}
