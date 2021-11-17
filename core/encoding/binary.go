package encoding

import (
	"bytes"
	"encoding/binary"

	"github.com/aceaura/libra/core/magic"
)

type Binary struct{}

func init() {
	register(NewBinary())
}

func NewBinary() *Binary {
	return new(Binary)
}

func (b Binary) String() string {
	return magic.TypeName(b)
}

func (Binary) Marshal(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (Binary) Unmarshal(data []byte, v interface{}) error {
	buf := bytes.NewReader(data)
	return binary.Read(buf, binary.LittleEndian, v)
}

func (b Binary) Reverse() Encoding {
	return b
}

type LittleEndian struct{}

func init() {
	register(NewLittleEndian())
}

func NewLittleEndian() *LittleEndian {
	return new(LittleEndian)
}

func (le LittleEndian) String() string {
	return magic.TypeName(le)
}

func (LittleEndian) Marshal(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (LittleEndian) Unmarshal(data []byte, v interface{}) error {
	buf := bytes.NewReader(data)
	return binary.Read(buf, binary.LittleEndian, v)
}

func (le LittleEndian) Reverse() Encoding {
	return le
}

type BigEndian struct{}

func init() {
	register(NewBigEndian())
}

func NewBigEndian() *BigEndian {
	return new(BigEndian)
}

func (be BigEndian) String() string {
	return magic.TypeName(be)
}

func (BigEndian) Marshal(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (BigEndian) Unmarshal(data []byte, v interface{}) error {
	buf := bytes.NewReader(data)
	return binary.Read(buf, binary.BigEndian, v)
}

func (be BigEndian) Reverse() Encoding {
	return be
}