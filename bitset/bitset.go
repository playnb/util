package bitset

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

type Bitset interface {
	Set(index uint32)
	Clear(index uint32)
	ClearAll()
	Test(index uint32) bool
	Range() []uint32
}

//固定长度，可并发 集合
type bitset struct {
	data []uint64
}

func New(size uint32) Bitset {
	bs := &bitset{}
	bs.data = make([]uint64, size/64, size/64)
	return bs
}

func (bs *bitset) Set(index uint32) {
	if bs == nil {
		return
	}
	mask := uint64(1) << (index % 64)
	if index < uint32(len(bs.data)*64) {
		for {
			v := bs.data[index/64]
			if atomic.CompareAndSwapUint64(&bs.data[index/64], v, v|mask) {
				return
			}
			runtime.Gosched()
		}
		//bs.data[index/64] |= mask
		//return
	}
	panic(fmt.Sprintf("out of range(%d of %d)", index, len(bs.data)*64))
}

func (bs *bitset) Clear(index uint32) {
	if bs == nil {
		return
	}
	mask := ^(uint64(1) << (index % 64))
	if index < uint32(len(bs.data)*64) {
		for {
			v := bs.data[index/64]
			if atomic.CompareAndSwapUint64(&bs.data[index/64], v, v&mask) {
				return
			}
			runtime.Gosched()
		}
		//bs.data[index/64] &= mask
		return
	}
	panic(fmt.Sprintf("out of range(%d of %d)", index, len(bs.data)*64))
}

func (bs *bitset) ClearAll() {
	if bs == nil {
		return
	}
	for i := 0; i < len(bs.data); i++ {
		bs.data[i] = 0
	}
}

func (bs *bitset) Test(index uint32) bool {
	if bs == nil {
		return false
	}
	if index < uint32(len(bs.data)*64) {
		return bs.data[index/64]&(uint64(1)<<(index%64)) != 0
	}
	return false
}

func (bs *bitset) Range() []uint32 {
	if bs == nil {
		return nil
	}
	var ret []uint32
	for i := 0; i < len(bs.data)*64; i++ {
		if bs.Test(uint32(i)) {
			ret = append(ret, uint32(i))
		}
	}
	return ret
}
