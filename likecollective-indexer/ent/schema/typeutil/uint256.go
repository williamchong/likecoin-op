package typeutil

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/holiman/uint256"
)

// uint256 from ent -> pg bigint
// but pg bigint is signed and its max value is 9223372036854775807
// need to coerce it to pg numeral for max uint64 value 18446744073709551615
//
// pg bigint max:   9223372036854775807
// pg numeral max: 18446744073709551615
type Uint256 *uint256.Int

var Uint256Type = Uint256(uint256.NewInt(0))

var Uint256ValueScanner = field.ValueScannerFunc[Uint256, *sql.NullString]{
	V: func(s Uint256) (driver.Value, error) {
		if s == nil {
			return "0", nil
		}
		v := (*uint256.Int)(s).String()
		return v, nil
	},
	S: func(ns *sql.NullString) (Uint256, error) {
		if !ns.Valid {
			return Uint256(uint256.NewInt(0)), nil
		}
		_s, err := uint256.FromDecimal(ns.String)
		if err != nil {
			return Uint256(uint256.NewInt(0)), err
		}
		return Uint256(_s), nil
	},
}

var Uint256SchemaType = map[string]string{
	dialect.Postgres: "numeric",
}

func Uint256Annotations(field string) []schema.Annotation {
	return []schema.Annotation{
		&entsql.Annotation{
			// The `Check` option allows adding an
			// unnamed CHECK constraint to table DDL.
			Checks: map[string]string{
				fmt.Sprintf("uint256_%s_check", field): fmt.Sprintf("%s >= 0", field),
			},
			Default: "0",
		},
	}
}
