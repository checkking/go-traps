package go_traps

import (
	"testing"
)

func TestSlices(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	b := a[1:3]
	t.Logf("a=%v\n", a)
	t.Logf("b=%v\n", b)
	a[1] = 10
	t.Logf("a=%v\n", a)
	t.Logf("b=%v\n", b)
	b = append(b, 11)
	b = append(b, 12)
	b = append(b, 13)
	b = append(b, 14)
	b[1] = 33
	a[1] = 99
	t.Logf("a=%v\n", a)
	t.Logf("b=%v\n", b)
}
