package snowflake

import (
	"sync"
	"testing"
)

var _all = make(map[uint64]bool)
var _allMutex = &sync.Mutex{}
var _allWait = &sync.WaitGroup{}

func genUUID(sf *SnowFlake) {
	defer _allWait.Done()
	for i := 0; i < 100000; i++ {
		uuid, err := sf.Next()
		if err != nil {
			panic(err)
		}
		_allMutex.Lock()
		if _, ok := _all[uuid]; ok {
			panic("...")
		}
		_all[uuid] = true
		_allMutex.Unlock()
	}
}

func Test_SnowFlake(t *testing.T) {
	for i := 1; i < 1000; i++ {
		sf, err := New(uint64(i))
		if err != nil {
			panic(err)
		}
		_allWait.Add(1)
		go genUUID(sf)
	}
	_allWait.Wait()
}
