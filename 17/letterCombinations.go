package letterCombinations

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var ans []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	ans = []string{}
	backTrack(digits, 0, "")
	return ans
}

func backTrack(digits string, index int, combination string) {
	if index == len(digits) {
		ans = append(ans, combination)
		return
	}

	digit := string(digits[index])
	phone := phoneMap[digit]
	for i := 0; i < len(phone); i++ {
		backTrack(digits, index+1, combination+string(phone[i]))
	}
}
