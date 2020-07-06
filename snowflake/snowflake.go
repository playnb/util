package snowflake

import (
	"errors"
	"github.com/playnb/util"
	"sync"
	"time"
)

//TODO  学习一下 https://github.com/sony/sonyflake  https://github.com/bwmarrin/snowflake 程序的细节

// Twitter-Snowflake
// +-------------------------------------------------------+
// | 41 Bit Timestamp | 10 Bit NodeID | 12 Bit Sequence ID |
// +-------------------------------------------------------+

const (
	MaxNodeId    = uint64(1<<BitesNodeID) - 1
	MaxTimestamp = uint64(1<<BitesTimestamp) - 1
	MaxSequence  = uint64(1<<BitesSequence) - 1
	epoch        = int64(1577808000) //2020/01/01

	BitesTimestamp = 41
	BitesNodeID    = 10
	BitesSequence  = 12
)

func New(nodeId uint64) (*SnowFlake, error) {
	if nodeId > MaxNodeId {
		return nil, errors.New("invalid node Id")
	}
	sf := &SnowFlake{nodeId: nodeId}
	err := sf.newEpoch()
	if err != nil {
		return nil, err
	}
	return sf, nil
}

type SnowFlake struct {
	timestamp uint64
	sequence  uint64
	nodeId    uint64
	lock      sync.Mutex
}

func (sf *SnowFlake) newEpoch() error {
	sf.timestamp = uint64(util.NowTimestampMillisecond() - epoch*1000)
	if sf.timestamp >= MaxTimestamp {
		return errors.New("timestamp overflow")
	}
	sf.sequence = 0
	return nil
}

func (sf *SnowFlake) Next() (uint64, error) {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	sf.sequence++
	if sf.sequence >= MaxSequence {
		time.Sleep(5 * time.Microsecond)
		err := sf.newEpoch()
		if err != nil {
			return 0, err
		}
	}

	return (sf.timestamp << (BitesSequence + BitesNodeID)) | (uint64(sf.nodeId) << BitesSequence) | (uint64(sf.sequence)), nil
}
