package optional_test

import (
	"encoding/json"
	"testing"

	"github.com/hhk7734/optional.go"
	"github.com/stretchr/testify/assert"
)

func stringPtr(v string) *string {
	return &v
}

func TestNil(t *testing.T) {
	var sPtr optional.Optional[*string]
	assert.Nil(t, sPtr)
	sPtr = optional.New[*string](nil)
	assert.NotNil(t, sPtr)
}

func TestStringPtrMarshalJSON(t *testing.T) {
	type testStruct struct {
		Value optional.Optional[*string] `json:"value,omitempty"`
	}

	cases := []struct {
		in   testStruct
		want []byte
	}{
		{
			in:   testStruct{},
			want: []byte(`{}`),
		},
		{
			in: testStruct{
				Value: optional.New[*string](nil),
			},
			want: []byte(`{"value":null}`),
		},
		{
			in: testStruct{
				Value: optional.New(stringPtr("")),
			},
			want: []byte(`{"value":""}`),
		},
		{
			in: testStruct{
				Value: optional.New(stringPtr("test")),
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

func TestStringPtrUnmarshalJSON(t *testing.T) {
	type testStruct struct {
		Value optional.Optional[*string] `json:"value,omitempty"`
	}

	cases := []struct {
		in   []byte
		want testStruct
	}{
		{
			in:   []byte(`{}`),
			want: testStruct{},
		},
		{
			in: []byte(`{"value":null}`),
			want: testStruct{
				Value: optional.New[*string](nil),
			},
		},
		{
			in: []byte(`{"value":""}`),
			want: testStruct{
				Value: optional.New(stringPtr("")),
			},
		},
		{
			in: []byte(`{"value":"test"}`),
			want: testStruct{
				Value: optional.New(stringPtr("test")),
			},
		},
	}

	for _, c := range cases {
		var got testStruct
		err := json.Unmarshal(c.in, &got)
		if assert.NoError(t, err) {
			assert.Equal(t, c.want, got)
		}
	}
}

func TestStringUnmarshalJSON(t *testing.T) {
	type testStruct struct {
		Value optional.Optional[string] `json:"value,omitempty"`
	}

	cases := []struct {
		in   []byte
		want testStruct
	}{
		{
			in:   []byte(`{}`),
			want: testStruct{},
		},
		{
			in: []byte(`{"value":""}`),
			want: testStruct{
				Value: optional.New(""),
			},
		},
		{
			in: []byte(`{"value":"test"}`),
			want: testStruct{
				Value: optional.New("test"),
			},
		},
	}

	for _, c := range cases {
		var got testStruct
		err := json.Unmarshal(c.in, &got)
		if assert.NoError(t, err) {
			assert.Equal(t, c.want, got)
		}
	}
}

func TestStringMarshalJSON(t *testing.T) {
	type testStruct struct {
		Value optional.Optional[string] `json:"value,omitempty"`
	}

	cases := []struct {
		in   testStruct
		want []byte
	}{
		{
			in:   testStruct{},
			want: []byte(`{}`),
		},
		{
			in: testStruct{
				Value: optional.New(""),
			},
			want: []byte(`{"value":""}`),
		},
		{
			in: testStruct{
				Value: optional.New("test"),
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
