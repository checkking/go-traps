package go_traps

import (
	"fmt"
	"testing"
)

func TestForRange(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	b := make([]*int, 0)

	for _, v := range a {
		b = append(b, &v)
	}
	for _, v := range b {
		fmt.Println(*v)
	}
}

