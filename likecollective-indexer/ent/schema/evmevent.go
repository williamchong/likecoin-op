package schema

import (
	"likecollective-indexer/ent/schema/typeutil"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// EVMEvent holds the schema definition for the EVMEvent entity.
type EVMEvent struct {
	ent.Schema
}

// Fields of the EVMEvent.
func (EVMEvent) Fields() []ent.Field {
	return []ent.Field{
		// Multiple events may share the same transaction hash
		// E.g. mint multiple
		field.String("transaction_hash").
			NotEmpty().
			Immutable(),
		field.Uint("transaction_index").
			Immutable(),
		field.Uint64("chain_id").
			GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner).
			Annotations(typeutil.Uint64Annotations("chain_id")...).
			Immutable(),
		field.String("block_hash").
			NotEmpty().
			Immutable(),
		field.Uint64("block_number").
			GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner).
			Annotations(typeutil.Uint64Annotations("block_number")...).
			Immutable(),
		field.Uint("log_index").
			Immutable(),
		field.String("address").
			NotEmpty().
			Immutable(),
		field.String("topic0").
			NotEmpty().
			Immutable(),
		field.String("topic0_hex").
			NotEmpty().
			Immutable(),
		field.String("topic1").
			Nillable().
			Optional().
			Immutable(),
		field.String("topic1_hex").
			Nillable().
			Optional().
			Immutable(),
		field.String("topic2").
			Nillable().
			Optional().
			Immutable(),
		field.String("topic2_hex").
			Nillable().
			Optional().
			Immutable(),
		field.String("topic3").
			Nillable().
			Optional().
			Immutable(),
		field.String("topic3_hex").
			Nillable().
			Optional().
			Immutable(),
		field.String("data").
			Nillable().
			Optional().
			Immutable(),
		field.String("data_hex").
			Nillable().
			Optional().
			Immutable(),
		field.Bool("removed").
			Immutable(),
		field.Enum("status").Values(
			"received",
			"skipped",
			"enqueued",
			"processing",
			"processed",
			"failed",
		),
		field.String("name").
			Immutable(),
		field.String("signature").
			Immutable(),
		field.JSON("indexed_params", map[string]any{}).
			Immutable(),
		field.JSON("non_indexed_params", map[string]any{}).
			Immutable(),
		field.String("failed_reason").
			Nillable().
			Optional().
			Immutable(),
		field.Time("timestamp").
			Immutable(),
	}
}

// Edges of the EVMEvent.
func (EVMEvent) Edges() []ent.Edge {
	return nil
}

func (EVMEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(
			"transaction_hash",
			"transaction_index",
			"block_number",
			"log_index",
		).Unique(),
		index.Fields("block_number"),
		index.Fields("log_index"),
		index.Fields("address"),
		index.Fields("signature"),
	}
}
