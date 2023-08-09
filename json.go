package optional

import (
	"bytes"
	"encoding/json"
	"reflect"
)

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o == nil {
		return []byte("null"), nil
	}

	marshal, err := json.Marshal(o[valueKey])
	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if len(data) <= 0 {
		*o = nil
		return nil
	}

	if bytes.Equal(data, []byte("null")) {
		var v T
		if reflect.ValueOf(v).Kind() == reflect.Ptr {
			*o = New(v)
		} else {
			*o = nil
		}
		return nil
	}

	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	*o = New(v)

	return nil
}
