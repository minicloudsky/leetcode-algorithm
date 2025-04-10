package solutions

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	stack := make([]byte, 0)
	sm := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < n; i++ {
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) == 0 || (stack[len(stack)-1] != sm[s[i]]) {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
