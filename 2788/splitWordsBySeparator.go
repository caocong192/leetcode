package splitWordsBySeparator

import "strings"

//func splitWordsBySeparator(words []string, separator byte) (ans []string) {
//	for _, w := range words {
//		if strings.Contains(w, string(separator)) {
//			res := ""
//			for _, v := range w {
//				if string(v) == string(separator) {
//					if res != "" {
//						ans = append(ans, res)
//						res = ""
//					}
//				} else {
//					res = res + string(v)
//				}
//			}
//
//			if res != "" {
//				ans = append(ans, res)
//			}
//		} else {
//			ans = append(ans, w)
//		}
//	}
//	return
//}

func splitWordsBySeparator(words []string, separator byte) (ans []string) {
	for _, word := range words {
		subs := strings.Split(word, string(separator))
		for _, sub := range subs {
			if len(sub) > 0 {
				ans = append(ans, sub)
			}
		}
	}
	return
}
