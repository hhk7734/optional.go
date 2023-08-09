package optional

type Optional[T any] []T

const valueKey = iota

func New[T any](value T) Optional[T] {
	return Optional[T]{
		valueKey: value,
	}
}

func (o Optional[T]) Get() (T, bool) {
	if o == nil {
		var v T
		return v, false
	}

	return o[valueKey], true
}

func (o Optional[T]) MustGet() T {
	v, _ := o.Get()
	return v
}
