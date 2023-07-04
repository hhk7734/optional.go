package optional_test

import (
	"encoding/json"
	"testing"

	"github.com/hhk7734/optional.go"
	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	type stringPointer struct {
		Value optional.Optional[*string] `json:"value,omitempty"`
	}

	testString := "test"

	cases := []struct {
		in   stringPointer
		want []byte
	}{
		{
			in:   stringPointer{},
			want: []byte(`{}`),
		},
		{
			in: stringPointer{
				Value: optional.New[*string](nil),
			},
			want: []byte(`{"value":null}`),
		},
		{
			in: stringPointer{
				Value: optional.New(&testString),
			},
			want: []byte(`{"value":"test"}`),
		},
	}

	for _, c := range cases {
		got, err := json.Marshal(c.in)
		if assert.NoError(t, err) {
			assert.Equal(t, c.want, got)
		}
	}

}

func TestUnmarshalJSON(t *testing.T) {
	type stringPointer struct {
		Value optional.Optional[*string] `json:"value,omitempty"`
	}

	testString := "test"

	cases := []struct {
		in   []byte
		want stringPointer
	}{
		{
			in:   []byte(`{}`),
			want: stringPointer{},
		},
		{
			in: []byte(`{"value":null}`),
			want: stringPointer{
				Value: optional.New[*string](nil),
			},
		},
		{
			in: []byte(`{"value":"test"}`),
			want: stringPointer{
				Value: optional.New(&testString),
			},
		},
	}

	for _, c := range cases {
		var got stringPointer
		err := json.Unmarshal(c.in, &got)
		if assert.NoError(t, err) {
			assert.Equal(t, c.want, got)
		}
	}
}
