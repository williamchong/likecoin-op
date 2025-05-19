package schema

import (
	"slices"

	"likenft-indexer/ent/schema/typeutil"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// NFT holds the schema definition for the NFT entity.
type LikeProtocol struct {
	ent.Schema
}

// Fields of the NFTClass.
func (LikeProtocol) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Unique(),
		field.Uint64("latest_event_block_number").GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
	}
}

func (LikeProtocol) Annotations() []schema.Annotation {
	return slices.Concat(
		typeutil.Uint64Annotations("latest_event_block_number"),
	)
}
