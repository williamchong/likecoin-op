package schema

import (
	"slices"

	"likenft-indexer/ent/schema/typeutil"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type TransactionMemo struct {
	ent.Schema
}

// Fields of the EVMEvent.
func (TransactionMemo) Fields() []ent.Field {
	return []ent.Field{
		field.String("transaction_hash").NotEmpty(),
		field.String("book_nft_id").NotEmpty().
			Comment("contract address of book nft"),
		field.String("from").NotEmpty(),
		field.String("to").NotEmpty(),
		field.Uint64("token_id").GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
		field.String("memo"),
		field.Uint64("block_number").GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
	}
}

// Edges of the EVMEvent.
func (TransactionMemo) Edges() []ent.Edge {
	return nil
}

func (TransactionMemo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("transaction_hash", "book_nft_id", "token_id").Unique(),
		index.Fields("book_nft_id", "token_id"),
	}
}

func (TransactionMemo) Annotations() []schema.Annotation {
	return slices.Concat(
		typeutil.Uint64Annotations("token_id"),
		typeutil.Uint64Annotations("block_number"),
	)
}
