package go_traps

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	var c chan int
	go func() {
		c <- 1
		t.Logf("send\n")
	}()
	fmt.Println("hello ")
	go func() {
		e := <- c
		t.Logf("world, e:%v ", e)
	}()
	time.Sleep(time.Second * 10)
}
