package jsondatauri

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type JSONDataUri string

func (s JSONDataUri) DataURIPrefix() string {
	return "data:application/json;utf8,"
}

func (s JSONDataUri) HttpPrefix() string {
	return "http"
}

func (s JSONDataUri) IsDataURIJson() (string, bool) {
	// Backward compatability for v0.1.0 and v0.2.0 which rewrites the metadata
	// from "data:application/json;utf8," to "data:application/json; charset=utf-8,"
	// for better block explorer compatability
	if strings.HasPrefix(string(s), "data:application/json;utf8,") {
		r, _ := strings.CutPrefix(string(s), "data:application/json;utf8,")
		return r, true
	}
	if strings.HasPrefix(string(s), "data:application/json; charset=utf-8,") {
		r, _ := strings.CutPrefix(string(s), "data:application/json; charset=utf-8,")
		return r, true
	}
	return "", false
}

func (s JSONDataUri) IsHttpJson() bool {
	return strings.HasPrefix(string(s), s.HttpPrefix())
}

func (s JSONDataUri) Resolve(httpClient HTTPClient, out any) error {
	err := json.Unmarshal([]byte(s), &out)

	// Can just be parsed
	if err == nil {
		return nil
	}

	if s.IsHttpJson() {
		req, err := http.NewRequest("GET", string(s), nil)
		if err != nil {
			return err
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&out)
		if err != nil {
			return err
		}
		return nil
	}

	r, yes := s.IsDataURIJson()
	if yes {
		return json.Unmarshal([]byte(r), &out)
	}

	return fmt.Errorf("unknwon string format %v", string(s))
}
