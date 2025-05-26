package handler

import (
	"net/http"
)

type PanicHandler struct{}

func (p *PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("panic from api")
}
