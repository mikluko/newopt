package newopt

func New[T any](opts ...Option[T]) (rv T) {
	for _, opt := range opts {
		rv = opt(rv)
	}
	return rv
}

func NewE[T any](opts ...OptionE[T]) (rv T, err error) {
	for _, opt := range opts {
		rv, err = opt(rv)
		if err != nil {
			return rv, err
		}
	}
	return rv, nil
}

func NewP[T any](opts ...Option[*T]) (rv *T) {
	rv = new(T)
	for _, opt := range opts {
		rv = opt(rv)
	}
	return rv
}

func NewPE[T any](opts ...OptionE[*T]) (rv *T, err error) {
	rv = new(T)
	for _, opt := range opts {
		rv, err = opt(rv)
		if err != nil {
			return rv, err
		}
	}
	return rv, nil
}

type Option[T any] func(T) T

type OptionE[T any] func(T) (T, error)
