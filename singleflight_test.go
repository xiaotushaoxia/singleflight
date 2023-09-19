package singleflight

import (
	"fmt"
	"testing"
	"time"
)

func TestGroup_Do(t *testing.T) {
	var g Group[int, int]
	for i := 0; i < 10; i++ {
		go func(ii int) {
			fmt.Println(g.Do(ii, func() (int, error) {
				return kk(ii)
			}))
		}(i / 2)
	}
	time.Sleep(time.Second * 5)
}

func kk(i int) (int, error) {
	time.Sleep(time.Second)
	fmt.Println("call kk ", i)
	return i, nil
}
