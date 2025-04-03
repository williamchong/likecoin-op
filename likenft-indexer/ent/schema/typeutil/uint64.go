package typeutil

import (
	"encoding"
	"fmt"
	"strconv"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// uint64 from ent -> pg bigint
// but pg bigint is signed and its max value is 9223372036854775807
// need to coerce it to pg numeral for max uint64 value 18446744073709551615
//
// pg bigint max:   9223372036854775807
// pg numeral max: 18446744073709551615
type Uint64 uint64

var _ encoding.TextMarshaler = Uint64(0)
var _ encoding.TextUnmarshaler = Uint64(0)

func (s Uint64) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(uint64(s), 10)), nil
}

func (s Uint64) UnmarshalText(text []byte) error {
	_s, err := strconv.ParseUint(string(text), 10, 64)
	if err != nil {
		return err
	}

	var sp *Uint64 = (*Uint64)(&s)
	*sp = Uint64(_s)
	return nil
}

var Uint64ValueScanner = &field.TextValueScanner[Uint64]{}

var Uint64SchemaType = map[string]string{
	dialect.Postgres: "numeric",
}

func Uint64Annotations(field string) []schema.Annotation {
	return []schema.Annotation{
		&entsql.Annotation{
			// The `Check` option allows adding an
			// unnamed CHECK constraint to table DDL.
			Checks: map[string]string{
				fmt.Sprintf("uint64_%s_check", field): fmt.Sprintf("%s >= 0", field),
			},
		},
	}
}
