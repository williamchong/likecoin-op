package schema

import (
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
		field.String("transaction_hash").NotEmpty().Unique(),
		field.String("block_hash").NotEmpty(),
		field.Uint64("block_number"),
		field.Uint64("log_index"),
		field.String("address").NotEmpty(),
		field.String("topic0").NotEmpty(),
		field.String("topic1"),
		field.String("topic2"),
		field.String("topic3"),
		field.String("data"),
		field.Time("timestamp"),
	}
}

// Edges of the EVMEvent.
func (EVMEvent) Edges() []ent.Edge {
	return nil
}

func (EVMEvent) Index() []ent.Index {
	return []ent.Index{
		index.Fields("block_number"),
		index.Fields("log_index"),
		index.Fields("address"),
	}
}
