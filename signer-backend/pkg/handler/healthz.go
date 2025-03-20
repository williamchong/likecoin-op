package handler

import (
	"fmt"
	"net/http"
)

type HealthzHandler struct{}

func MakeHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
