package atomicswap_test

import (
	"github.com/GodsBoss/atomicswap"

	"fmt"
	"testing"
)

func TestSwapper(t *testing.T) {
	swapper := atomicswap.NewSwapper("foo")

	go func() {
		for i := 0; i < 100000; i++ {
			swapper.Set(fmt.Sprintf("%d", i))
		}
	}()

	for i := 0; i < 10000; i++ {
		go func() {
			for j := 0; j < 100000; j++ {
				_ = swapper.Get()
			}
		}()
	}
}
