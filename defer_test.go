package go_traps

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	fmt.Println("a return:", a())
}

func TestDefer2(t *testing.T) {
	fmt.Println("b return:", b())
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i)
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("a defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i)
	}()
	return
}
