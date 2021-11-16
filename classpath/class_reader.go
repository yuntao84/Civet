package classpath

import (
	"encoding/binary"
	"io"
)

type ClassReader struct {
	data io.Reader
}

func (self *ClassReader) ReadUint8() (uint8, error) {
	var val uint8
	err := binary.Read(self.data, binary.BigEndian, &val)
	return val, err
}

func (self *ClassReader) ReadUint16() (uint16, error) {
	var val uint16
	err := binary.Read(self.data, binary.BigEndian, &val)
	return val, err
}

func (self *ClassReader) ReadUint32() (uint32, error) {
	var val uint32
	err := binary.Read(self.data, binary.BigEndian, &val)
	return val, err
}

func (self *ClassReader) ReadUint64() (uint64, error) {
	var val uint64
	err := binary.Read(self.data, binary.BigEndian, &val)
	return val, err
}

func (self *ClassReader) ReadUint16s() ([]uint16, error) {
	n, err := self.ReadUint16()
	if err != nil {
		return nil, err
	}
	val := make([]uint16, n)
	err = binary.Read(self.data, binary.BigEndian, &val)
	return val, nil
}

func (self *ClassReader) ReadBytes(n uint32) ([]byte, error) {
	val := make([]byte, n)
	_, err := self.data.Read(val)
	return val, err
}
