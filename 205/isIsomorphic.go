package isIsomorphic

func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}

		s2t[x] = y
		t2s[y] = x
	}
	return true
}
