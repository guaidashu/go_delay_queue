/**
 * @Author: yy
 * @Author: 1023767856@qq.com
 * @Date: 28/06/2021
 * @Desc: desc
 */

package go_delay_queue

import "sync/atomic"

type AtomicBool uint32

// NewAtomicBool returns an AtomicBool.
func NewAtomicBool() *AtomicBool {
	return new(AtomicBool)
}

// Set sets the value to v.
func (b *AtomicBool) Set(v bool) {
	if v {
		atomic.StoreUint32((*uint32)(b), 1)
	} else {
		atomic.StoreUint32((*uint32)(b), 0)
	}
}

// True returns true if current value is true.
func (b *AtomicBool) True() bool {
	return atomic.LoadUint32((*uint32)(b)) == 1
}

// CompareAndSwap compares current value with given old, if equals, set to given val.
func (b *AtomicBool) CompareAndSwap(old, val bool) bool {
	var ov, nv uint32
	if old {
		ov = 1
	}
	if val {
		nv = 1
	}

	return atomic.CompareAndSwapUint32((*uint32)(b), ov, nv)
}

// ForAtomicBool returns an AtomicBool with given val.
func ForAtomicBool(val bool) *AtomicBool {
	b := NewAtomicBool()
	b.Set(val)
	return b
}