package countTestedDevices

func countTestedDevices(batteryPercentages []int) (ans int) {
	reduce := 1
	for i := 0; i < len(batteryPercentages); i++ {
		if batteryPercentages[i]-reduce >= 0 {
			reduce++
			ans++
		}
	}
	return
}
