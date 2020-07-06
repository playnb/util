package snowflake

import (
	c "github.com/smartystreets/goconvey/convey"
	"os"
	"sync"
	"testing"
)

var _all = make(map[uint64]bool)
var _allMutex = &sync.Mutex{}
var _allWait = &sync.WaitGroup{}

const (
	N = uint64(10)
	M = 5 * MaxSequence
	C = 3
)

func genUUID(sf *SnowFlake) {
	defer _allWait.Done()
	for i := uint64(0); i < M; i++ {
		uuid, err := sf.Next()
		c.So(err, c.ShouldBeNil)

		_allMutex.Lock()
		_, ok := _all[uuid]
		c.So(ok, c.ShouldBeFalse)
		_all[uuid] = true
		_allMutex.Unlock()

	}
}

func Test_SnowFlake(t *testing.T) {
	oldEnv := os.Getenv("GOCONVEY_REPORTER")
	defer func() {
		os.Setenv("GOCONVEY_REPORTER", oldEnv)
	}()
	os.Setenv("GOCONVEY_REPORTER", "silent")
	c.Convey("Test_SnowFlake", t, func() {
		_, err := New(MaxNodeId)
		c.ShouldBeError(err)

		for i := uint64(1); i < N; i++ {
			sf, err := New(i)
			sf.timestamp = 10000
			c.So(err, c.ShouldBeNil)
			for j := 0; j < C; j++ {
				_allWait.Add(1)
				go c.Convey("test one generator", t, func() { genUUID(sf) })
			}
		}
		_allWait.Wait()
		c.So(uint64(len(_all)), c.ShouldEqual, C*(N-1)*(M))
	})
}
