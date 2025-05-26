package types

import (
	"bytes"
	"encoding/json"
	"time"

	timeutil "github.com/likecoin/like-migration-backend/pkg/likenft/util/time"
)

type StringPreservedDateTime struct {
	str string
	dt  time.Time
}

func (t *StringPreservedDateTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`""`)) {
		t.str = ``
		return nil
	}
	dt, err := timeutil.TimeLayouts.Parse(string(bytes.Trim(data, "\"")))
	if err != nil {
		return err
	}
	var str string
	err = json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	t.str = str
	t.dt = dt
	return nil
}

func (t *StringPreservedDateTime) MarshalJSON() ([]byte, error) {
	if t.IsEmpty() {
		return []byte(`""`), nil
	}
	return []byte(t.str), nil
}

func (t *StringPreservedDateTime) GetEpochSeconds() int64 {
	return t.dt.Unix()
}

func (t *StringPreservedDateTime) IsEmpty() bool {
	return t.str == ""
}

func (t *StringPreservedDateTime) ToString() string {
	return t.str
}
