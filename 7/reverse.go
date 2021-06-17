package reverse

import (
	"fmt"
)

func reverse(x int) int {
	strX := fmt.Sprintf("%v", x)
	minus := false

	if string(strX[0]) == "-" {
		minus = true
	}

	if minus && len(strX) > 11 || !minus && len(strX) > 10 {
		return 0
	}else if

	edge := 2187483647

}
