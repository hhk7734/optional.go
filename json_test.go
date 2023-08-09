package optional_test

import (
	"encoding/json"
	"testing"

	"github.com/hhk7734/optional.go"
	"github.com/stretchr/testify/assert"
)

func TestStringPtrMarshalJSON(t *testing.T) {
	type Struct struct {
		Field optional.Optional[*string] `json:"field,omitempty"`
	}

	cases := []struct {
		in   Struct
		want []byte
	}{
		{
			in:   Struct{},
			want: []byte(`{}`),
		},
		{
			in: Struct{
				Field: optional.New[*string](nil),
			},
			want: []byte(`{"field":null}`),
		},
		{
			in: Struct{
				Field: optional.New(stringPtr("")),
			},
			want: []byte(`{"field":""}`),
		},
		{
			in: Struct{
				Field: optional.New(stringPtr("test")),
			},
			want: []byte(`{"field":"test"}`),
		},
	}

	for _, c := range cases {
		got, err := json.Marshal(c.in)
		if assert.NoError(t, err) {
			assert.Equal(t, c.want, got)
		}
	}
}

func TestStringPtrUnmarshalJSON(t *testing.T) {
	type Struct struct {
		Field optional.Optional[*string] `json:"field,omitempty"`
	}

	cases := []struct {
		in   []byte
		want Struct
	}{
		{
			in:   []byte(`{}`),
			want: Struct{},
		},
		{
			in: []byte(`{"field":null}`),
			want: Struct{
				Field: optional.New[*string](nil),
			},
		},
		{
			in: []byte(`{"field":""}`),
			want: Struct{
				Field: optional.New(stringPtr("")),
			},
		},
		{
			in: []byte(`{"field":"test"}`),
			want: Struct{
				Field: optional.New(stringPtr("test")),
			},
		},
	}

	for _, c := range cases {
		var got Struct
		err := json.Unmarshal(c.in, &got)
		if assert.NoError(t, err) {
			assert.Equal(t, c.want, got)
		}
	}
}

func TestStringUnmarshalJSON(t *testing.T) {
	type Struct struct {
		Field optional.Optional[string] `json:"field,omitempty"`
	}

	cases := []struct {
		in   []byte
		want Struct
	}{
		{
			in:   []byte(`{}`),
			want: Struct{},
		},
		{
			in: []byte(`{"field":""}`),
			want: Struct{
				Field: optional.New(""),
			},
		},
		{
			in: []byte(`{"field":"test"}`),
			want: Struct{
				Field: optional.New("test"),
			},
		},
	}

	for _, c := range cases {
		var got Struct
		err := json.Unmarshal(c.in, &got)
		if assert.NoError(t, err) {
			assert.Equal(t, c.want, got)
		}
	}
}

func TestStringMarshalJSON(t *testing.T) {
	type Struct struct {
		Field optional.Optional[string] `json:"field,omitempty"`
	}

	cases := []struct {
		in   Struct
		want []byte
	}{
		{
			in:   Struct{},
			want: []byte(`{}`),
		},
		{
			in: Struct{
				Field: optional.New(""),
			},
			want: []byte(`{"field":""}`),
		},
		{
			in: Struct{
				Field: optional.New("test"),
			},
			want: []byte(`{"field":"test"}`),
		},
	}

	for _, c := range cases {
		got, err := json.Marshal(c.in)
		if assert.NoError(t, err) {
			assert.Equal(t, c.want, got)
		}
	}
}
