package schema

import (
	"slices"

	"likenft-indexer/ent/schema/typeutil"

	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("transaction_hash").NotEmpty(),
		field.Uint("transaction_index"),
		field.Uint64("chain_id").GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
		field.String("block_hash").NotEmpty(),
		field.Uint64("block_number").GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
		field.Uint("log_index"),
		field.String("address").NotEmpty(),
		field.String("topic0").NotEmpty(),
		field.String("topic0_hex").NotEmpty(),
		field.String("topic1").Nillable().Optional(),
		field.String("topic1_hex").Nillable().Optional(),
		field.String("topic2").Nillable().Optional(),
		field.String("topic2_hex").Nillable().Optional(),
		field.String("topic3").Nillable().Optional(),
		field.String("topic3_hex").Nillable().Optional(),
		field.String("data").Nillable().Optional(),
		field.String("data_hex").Nillable().Optional(),
		field.Bool("removed"),
		field.Enum("status").Values(
			"received",
			"enqueued",
			"processing",
			"processed",
			"failed",
		),
		field.String("name"),
		field.String("signature"),
		field.JSON("indexed_params", map[string]any{}),
		field.JSON("non_indexed_params", map[string]any{}),
		field.String("failed_reason").Nillable().Optional(),
		field.Time("timestamp"),
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

func (EVMEvent) Annotations() []schema.Annotation {
	return slices.Concat(
		typeutil.Uint64Annotations("block_number"),
	)
}
