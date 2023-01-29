package main

func countAsterisks(s string) (ans int) {
	count := 0
	for _, v := range s {
		if count%2 == 0 {
			if string(v) == "*" {
				ans++
			}
		}

		if string(v) == "|" {
			count++
		}
	}
	return
}


func main() {
	s := "yo|uar|e**|b|e***au|tifu|l"
	println(countAsterisks(s))
}
