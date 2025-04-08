package typeutil

import (
	"database/sql"
	"database/sql/driver"
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

var Uint64ValueScanner = field.ValueScannerFunc[Uint64, *sql.NullString]{
	V: func(s Uint64) (driver.Value, error) {
		return []byte(strconv.FormatUint(uint64(s), 10)), nil
	},
	S: func(ns *sql.NullString) (Uint64, error) {
		if !ns.Valid {
			return Uint64(0), nil
		}
		_s, err := strconv.ParseUint(string(ns.String), 10, 64)
		if err != nil {
			return Uint64(0), err
		}
		return Uint64(_s), nil
	},
}

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
