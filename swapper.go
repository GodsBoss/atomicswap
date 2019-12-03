package atomicswap

import (
	"sync/atomic"
)

type Swapper struct {
	index  *int64
	values map[int64]*value
}

func NewSwapper(current string) *Swapper {
	return &Swapper{
		index: new(int64),
		values: map[int64]*value{
			0: &value{
				s: current,
			},
			1: &value{
				s: "",
			},
		},
	}
}

func (swapper *Swapper) Get() string {
	return swapper.values[*swapper.index].s
}

func (swapper *Swapper) Set(s string) {
	otherIndex := 1 - *swapper.index
	swapper.values[otherIndex].set(s)
	swapped := atomic.CompareAndSwapInt64(swapper.index, *swapper.index, otherIndex)
	if !swapped {
		panic("not swapped!")
	}
}

type value struct {
	s string
}

func (v *value) set(s string) {
	v.s = s
}
