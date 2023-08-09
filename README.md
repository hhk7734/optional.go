## optional.go

### optional.Optional <-\> JSON

```go
type Struct struct {
    Field optional.Optional[*string] `json:"field,omitempty"`
}
```

```go
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
```

```go
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
```