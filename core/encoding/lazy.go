package encoding

import (
	"errors"

	"github.com/aceaura/libra/magic"
)

var (
	ErrLazyWrongValueType = errors.New("codec lazy converts on wrong type value")
)

type Lazy struct{}

func init() {
	register(NewLazy())
}

func NewLazy() *Lazy {
	return new(Lazy)
}

func (lazy Lazy) String() string {
	return magic.TypeName(lazy)
}

func (Lazy) Marshal(v interface{}) ([]byte, error) {
	switch v := v.(type) {
	case []byte:
		data := make([]byte, len(v))
		copy(data, v)
		return data, nil
	case Bytes:
		data := make([]byte, len(v.Data))
		copy(data, v.Data)
		return data, nil
	case *Bytes:
		data := make([]byte, len(v.Data))
		copy(data, v.Data)
		return data, nil
	default:
		return nil, ErrLazyWrongValueType
	}
}

func (Lazy) Unmarshal(data []byte, v interface{}) error {
	switch v := v.(type) {
	case *Bytes:
		v.Data = make([]byte, len(data))
		copy(v.Data, data)
		return nil
	default:
		return ErrLazyWrongValueType
	}
}