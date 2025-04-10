package solutions

import (
	"fmt"
	"testing"
)

func TestIsVaild(t *testing.T) {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
