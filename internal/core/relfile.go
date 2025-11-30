package core

import (
	"encoding/binary"
	"errors"
	"os"
)

type Slot struct {
	Key    [16]byte
	Orig   int32
	Final  int32
	Probes int32
	Filled uint8
	_pad   [7]byte
}

const TableSize = 100
const SlotSize = 32

func Open(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	if info, _ := f.Stat(); info.Size() < TableSize*SlotSize {
		f.Truncate(TableSize * SlotSize)
	}
	return f, nil
}

func ReadSlot(f *os.File, idx int) (Slot, error) {
	var s Slot
	if idx < 0 || idx >= TableSize {
		return s, errors.New("bad index")
	}
	buf := make([]byte, SlotSize)
	_, err := f.ReadAt(buf, int64(idx*SlotSize))
	if err != nil {
		return s, err
	}
	copy(s.Key[:], buf[:16])
	s.Orig = int32(binary.LittleEndian.Uint32(buf[16:20]))
	s.Final = int32(binary.LittleEndian.Uint32(buf[20:24]))
	s.Probes = int32(binary.LittleEndian.Uint32(buf[24:28]))
	s.Filled = buf[28]
	return s, nil
}

func WriteSlot(f *os.File, idx int, s Slot) error {
	buf := make([]byte, SlotSize)
	copy(buf[:16], s.Key[:])
	binary.LittleEndian.PutUint32(buf[16:20], uint32(s.Orig))
	binary.LittleEndian.PutUint32(buf[20:24], uint32(s.Final))
	binary.LittleEndian.PutUint32(buf[24:28], uint32(s.Probes))
	buf[28] = s.Filled
	_, err := f.WriteAt(buf, int64(idx*SlotSize))
	return err
}
