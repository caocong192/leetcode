package discountPrices

import (
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	discountLists := strings.Split(sentence, " ")
	for i := 0; i < len(discountLists); i++ {
		if isPrice(discountLists[i]) {

			price, _ := strconv.ParseFloat(discountLists[i][1:], 64)
			discountLists[i] = "$" + strconv.FormatFloat(price*float64(100-discount)/100, 'f', 2, 64)
		}
	}
	return strings.Join(discountLists, " ")
}

func isPrice(str string) bool {
	if len(str) == 1 || str[0] != '$' {
		return false
	}

	countDot := 0
	for i := 1; i < len(str); i++ {
		if str[i] == '.' {
			countDot++
		}
		if str[i] < '0' || str[i] > '9' {
			return false
		}
	}
	if countDot > 1 {
		return false
	}
	return true
}
