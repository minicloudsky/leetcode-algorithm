package solutions

import "sort"

// GroupAnagrams https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
// 字母异位词分组
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//
// 示例 1:
//
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
// 示例 2:
//
// 输入: strs = [""]
// 输出: [[""]]
// 示例 3:
//
// 输入: strs = ["a"]
// 输出: [["a"]]
//
// 提示：
//
// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
// 相同字母排列得到的单词，放到一个数组中，通过排序实现
func GroupAnagrams(strs []string) [][]string {
	wordMap := make(map[string][]string)
	for _, s := range strs {
		charArr := []byte(s)
		sort.Slice(charArr, func(i, j int) bool {
			return charArr[i] < charArr[j]
		})
		wordMap[string(charArr)] = append(wordMap[string(charArr)], s)
	}
	res := make([][]string, 0)
	for _, w := range wordMap {
		res = append(res, w)
	}
	return res
}
