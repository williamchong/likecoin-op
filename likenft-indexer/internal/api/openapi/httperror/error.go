package httperror

import "errors"

// 401
var ErrUnauthorized = errors.New("unauthroized")
