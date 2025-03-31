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

func (s JSONDataUri) IsDataURIJson() bool {
	return strings.HasPrefix(string(s), s.DataURIPrefix())
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

	if s.IsDataURIJson() {
		r, _ := strings.CutPrefix(string(s), s.DataURIPrefix())
		return json.Unmarshal([]byte(r), &out)
	}

	return fmt.Errorf("unknwon string format %v", string(s))
}
