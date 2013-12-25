package worker

import (
	"encoding/binary"
)

// Worker side job
type outPack struct {
	dataType uint32
	data     []byte
	handle   string
}

func getOutPack() (outpack *outPack) {
	// TODO pool
	return &outPack{}
}

// Encode a job to byte slice
func (outpack *outPack) Encode() (data []byte) {
	var l int
	if outpack.dataType == WORK_FAIL {
		l = len(outpack.handle)
	} else {
		l = len(outpack.data)
		if outpack.handle != "" {
			l += len(outpack.handle) + 1
		}
	}
	data = getBuffer(l + MIN_PACKET_LEN)
	binary.BigEndian.PutUint32(data[:4], REQ)
	binary.BigEndian.PutUint32(data[4:8], outpack.dataType)
	binary.BigEndian.PutUint32(data[8:MIN_PACKET_LEN], uint32(l))
	i := MIN_PACKET_LEN
	if outpack.handle != "" {
		hi := len(outpack.handle) + i
		copy(data[i:hi], []byte(outpack.handle))
		if outpack.dataType != WORK_FAIL {
			data[hi] = '\x00'
		}
		i = hi + 1
	}
	if outpack.dataType != WORK_FAIL {
		copy(data[i:], outpack.data)
	}
	return
}