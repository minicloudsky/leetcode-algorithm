package string_match

// NaiveStringSearch 最简单直接的字符串匹配算法，逐个字符比较，时间复杂度为O(n*m)
func NaiveStringSearch(text, pattern string) int {
	n := len(text)
	m := len(pattern)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++
		}
		if j == m {
			return i // 返回匹配的起始位置
		}
	}
	return -1 // 没有找到匹配
}

func kMPComputeLPSArray(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	j := 0

	for i := 1; i < m; {
		if pattern[i] == pattern[j] {
			lps[i] = j + 1
			i++
			j++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// KMPSearch 利用已匹配的信息来避免不必要的字符比较，时间复杂度为O(n+m)。
func KMPSearch(text, pattern string) int {
	n := len(text)
	m := len(pattern)

	lps := kMPComputeLPSArray(pattern)
	i, j := 0, 0

	for i < n {
		if text[i] == pattern[j] {
			i++
			j++
		}
		if j == m {
			return i - j // 返回匹配的起始位置
		}
		if i < n && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1 // 没有找到匹配
}

func computeBadCharacterTable(pattern string) map[byte]int {
	table := make(map[byte]int)
	m := len(pattern)
	for i := 0; i < m; i++ {
		table[pattern[i]] = i
	}
	return table
}

func BMStringSearch(text, pattern string) int {
	n := len(text)
	m := len(pattern)
	badCharacterTable := computeBadCharacterTable(pattern)

	i := 0
	for i <= n-m {
		j := m - 1
		for j >= 0 && pattern[j] == text[i+j] {
			j--
		}
		if j < 0 {
			return i // 返回匹配的起始位置
		} else {
			skip := j - badCharacterTable[text[i+j]]
			if skip < 1 {
				skip = 1
			}
			i += skip
		}
	}
	return -1 // 没有找到匹配
}
