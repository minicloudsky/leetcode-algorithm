package pascals_triangle

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	num := 10
	ans := Generate(num)
	fmt.Println(ans)
}
