package model

import "likecollective-indexer/openapi/api"

type Opt[T any] interface {
	Get() (T, bool)
}

func FromOpt[T any](opt Opt[T]) *T {
	if v, ok := opt.Get(); ok {
		return &v
	}
	return nil
}

func NewOptString(s *string) api.OptString {
	if s == nil {
		return api.OptString{
			Value: "",
			Set:   false,
		}
	}
	return api.NewOptString(*s)
}
