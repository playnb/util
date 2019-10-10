package util

import (
	"fmt"
	"math"
	"sync"
)

var _defaultPool *BuffDataPool

func init() {
	_defaultPool = &BuffDataPool{}
	_defaultPool.Init()
}

func DefaultPool() *BuffDataPool {
	return _defaultPool
}

func makeBuffData(size int) *buffData {
	data := &buffData{}
	data.buffCap = size
	if size > 0 {
		data.data = make([]byte, size)
	}
	data.size = size
	data.index = 0
	data.inUse = true
	return data
}
func MakeBuffData(size int) BuffData {
	return makeBuffData(size)
}

func makeBuffDataBySlice(buf []byte, offset int) *buffData {
	data := makeBuffData(0)
	if offset > 0 {
		data.data = make([]byte, len(buf)+offset)
		data.index = 0
		data.size = len(data.data)
		data.ChangeIndex(offset)
		copy(data.data[data.index:], buf)
	} else {
		data.data = buf
		data.index = 0
		data.size = len(data.data)
	}
	return data
}
func MakeBuffDataBySlice(buf []byte, offset int) BuffData {
	return makeBuffDataBySlice(buf, offset)
}

type BuffData interface {
	GetPayload() []byte //只有数据部分
	Data() []byte       //包含index部分和数据部分
	Release()
	Size() int
	SetSize(size int)
	ChangeIndex(index int) bool
	Append(data []byte)
}

type buffData struct {
	buffCap int
	inUse   bool
	pool    *BuffDataPool

	size  int
	index int
	data  []byte
}

func (buff *buffData) equal(other *buffData) bool {
	if buff == nil && other == nil {
		return true
	}
	if buff == nil || other == nil {
		return false
	}
	a := buff.GetPayload()
	b := other.GetPayload()
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}
func (buff *buffData) Equal(other BuffData) bool {
	return buff.equal(other.(*buffData))
}

func (buff *buffData) GetPayload() []byte {
	if buff != nil {
		if buff.index+buff.size > len(buff.data) {
			//panic(buff)
			fmt.Printf("buffData 数据异常 %v", buff)
		}
		d := buff.data[buff.index : buff.index+buff.size]
		return d
	}
	return nil
}

func (buff *buffData) Release() {
	if buff != nil && buff.pool != nil {
		buff.pool.Put(buff)
	}
}

func (buff *buffData) Size() int {
	if buff != nil {
		return buff.size
	}
	return 0
}

func (buff *buffData) SetSize(size int) {
	if buff != nil {
		if buff.size < size {
			fmt.Println("buffData 设置大小出错")
			return
		}
		buff.size = size
	}
}

func (buff *buffData) Data() []byte {
	if buff != nil {
		return buff.data[:buff.index+buff.size]
	}
	return nil
}

func (buff *buffData) ChangeIndex(index int) bool {
	if buff != nil {
		if index > 0 && buff.size >= index {
			buff.index += index
			buff.size -= index
			return true
		} else if index < 0 && buff.index >= -index {
			buff.index += index
			buff.size -= index
			return true
		}
	}
	return false
}

func (buff *buffData) Append(data []byte) {
	if buff != nil {
		copy(buff.data[buff.index+buff.size:], data)
		buff.size = buff.size + len(data)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////
type OneBuffDataPool struct {
	sync.Pool
	cap int
}

func (pool *OneBuffDataPool) String() string {
	return ""
}

/*
type OneBuffDataPool struct {
	//sync.Pool
	cap       int
	New       func() interface{}
	buff      []interface{}
	buffIndex int
	lock      sync.Mutex
}

func (pool *OneBuffDataPool) String() string {
	pool.lock.Lock()
	defer pool.lock.Unlock()
	return fmt.Sprintf("缓存:%d Index:%d", len(pool.buff), pool.buffIndex)
}

func (pool *OneBuffDataPool) Get() interface{} {
	pool.lock.Lock()
	defer pool.lock.Unlock()

	if pool.buffIndex > 0 && len(pool.buff) > 0 {
		b := pool.buff[pool.buffIndex-1]
		pool.buffIndex--
		return b
	} else {
		return pool.New()
	}
	return nil
}

func (pool *OneBuffDataPool) Put(d interface{}) {
	pool.lock.Lock()
	defer pool.lock.Unlock()
	if pool.buffIndex < len(pool.buff) {
		pool.buffIndex++
		pool.buff[pool.buffIndex-1] = d
	} else {
		pool.buffIndex++
		pool.buff = append(pool.buff, d)
	}
}
*/
/////////////////////////////////////////////////////////////////////////////////////////////

type BuffDataPool struct {
	pools []*OneBuffDataPool
	lock  sync.Mutex
}

func (bp *BuffDataPool) Dump() string {
	str := ""
	for _, p := range bp.pools {
		if p != nil {
			str = fmt.Sprintf("%s%d:%s\n", str, p.cap, p.String())
		}
	}
	return str
}

func (bp *BuffDataPool) Init() {
	bp.lock.Lock()
	defer bp.lock.Unlock()
	bp.pools = make([]*OneBuffDataPool, 33)
	for i := uint32(3); i < 33; i++ {
		pool := &OneBuffDataPool{}
		pool.cap = int(uint32(1<<i) - 1)
		pool.New = func() interface{} {
			data := makeBuffData(pool.cap)
			data.pool = bp
			data.inUse = false
			return data
		}
		bp.pools[i] = pool
	}
}

func (bp *BuffDataPool) getPool(size int) *OneBuffDataPool {
	for i := 0; i < len(bp.pools); i++ {
		pool := bp.pools[i]
		if pool == nil {
			continue
		}
		if size <= pool.cap {
			return pool
		}
	}
	return nil
}

func (bp *BuffDataPool) Get(size int) BuffData {
	if size >= math.MaxInt32 {
		return MakeBuffData(size)
	}

	bp.lock.Lock()
	pool := bp.getPool(size)
	bp.lock.Unlock()

	if pool != nil {
		data := pool.Get().(*buffData)
		data.inUse = true
		data.index = 0
		data.size = size
		data.pool = bp
		return data
	}
	return nil
}

func (bp *BuffDataPool) Put(data BuffData) {
	bp.put(data.(*buffData))
}

func (bp *BuffDataPool) put(data *buffData) {
	if data == nil || data.pool != bp {
		return
	}

	bp.lock.Lock()
	pool := bp.getPool(data.buffCap)
	if data.inUse == true {
		data.inUse = false
	} else {
		pool = nil
	}
	bp.lock.Unlock()

	if pool != nil {
		pool.Put(data)
	} else {
		//fmt.Printf("BuffDataPool 不能归还的数据 %v\n", data)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////
//用于goconvey判断BuffData相同
func ShouldBuffDataEqual(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 {
		return "expected不能为空"
	}
	a, _ := actual.(*buffData)
	b, _ := expected[0].(*buffData)
	if a.equal(b) {
		return ""
	}
	return "BuffData不相同"
}

func ShouldByteSilceEqual(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 {
		return "expected不能为空"
	}
	a, _ := actual.([]byte)
	b, _ := expected[0].([]byte)
	if a == nil || b == nil {
		return ""
	}
	if len(a) != len(b) {
		return "[]byte长度不符"
	}

	for k, v := range a {
		if v != b[k] {
			return "[]byte内容不符"
		}
	}

	return ""
}
