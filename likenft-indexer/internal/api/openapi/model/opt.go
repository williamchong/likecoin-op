package model

type Opt[T any] interface {
	Get() (T, bool)
}

func FromOpt[T any](opt Opt[T]) *T {
	if v, ok := opt.Get(); ok {
		return &v
	}
	return nil
}
