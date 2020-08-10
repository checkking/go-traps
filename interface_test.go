package go_traps

import (
	"fmt"
	"testing"
)

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func TestInterface(t *testing.T) {
	var x *int = nil
	Foo(x)
}

