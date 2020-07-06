package util

import "sort"

//有时效性的数组

type TmpElement struct {
	Data      interface{}
	timestamp int64
}

type TmpSet struct {
	data         map[uint64]*TmpElement
	validityTime int64
}

func NewTmpSet(t int64) *TmpSet {
	return &TmpSet{
		data:         make(map[uint64]*TmpElement),
		validityTime: t,
	}
}

func (t *TmpSet) Insert(uid uint64, data interface{}) {
	e, ok := t.data[uid]
	if !ok {
		e = &TmpElement{
			Data:      data,
			timestamp: NowTimestamp(),
		}
		t.data[uid] = e
	}
	e.timestamp = NowTimestamp()
	e.Data = data

	sort.Slice(t.data, func(i, j int) bool {
		return false
		//return t.data[i].timestamp < t.data[j].timestamp
	})
}
