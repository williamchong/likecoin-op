package httputil

import (
	"errors"
	"net/http"
)

func HandleResponseStatus(
	resp *http.Response,
) error {
	if resp.StatusCode >= 400 {
		return errors.New(resp.Status)
	}
	return nil
}
