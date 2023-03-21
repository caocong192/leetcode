package convertTemperature

func convertTemperature(celsius float64) (ans []float64) {
	ans = append(ans, CtoK(celsius))
	ans = append(ans, CtoF(celsius))
	return
}

func CtoF(a float64) float64 {
	return a/5*9 + 32
}

func CtoK(a float64) float64 {
	return 273.15 + a
}
