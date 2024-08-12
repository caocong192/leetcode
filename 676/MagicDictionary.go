package MagicDictionary

type MagicDictionary struct {
	dict []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.dict = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, word := range this.dict {
		if len(word) != len(searchWord) {
			continue
		}

		if word == searchWord {
			continue
		}
		cnt := 0
		for i, ch := range word {
			if ch != rune(searchWord[i]) {
				cnt++
			}
			if cnt > 1 {
				break
			}
		}
		if cnt == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
