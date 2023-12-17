package jewels_and_stones

import (
	"fmt"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	jewels := "aA"
	stones := "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels, stones))
	jewels2 := "z"
	stones2 := "ZZ"
	fmt.Println(numJewelsInStones(jewels2, stones2))
}
