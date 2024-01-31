package string_match

// zhangsan
// san
func StringMatch(str, substr string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != substr[0] {
			continue
		}
		for j := 0; i < len(substr); j++ {
			if str[i] != substr[j] {
				break
			}
		}
	}
	return true
}
