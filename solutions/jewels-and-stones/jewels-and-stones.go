package jewels_and_stones

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[string]int, 0)
	for _, i := range jewels {
		jewelsMap[string(i)] = 0
	}
	number := 0
	for _, i := range stones {
		if _, ok := jewelsMap[string(i)]; ok {
			number += 1
		}
	}

	return number
}

func numJewelsInStones2(jewels string, stones string) int {
	number := 0
	for _, j := range jewels {
		for _, s := range stones {
			if j == s {
				number++
			}
		}
	}
	return number
}
