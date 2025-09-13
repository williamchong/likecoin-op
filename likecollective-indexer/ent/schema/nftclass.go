package schema

import (
	"likecollective-indexer/ent/schema/typeutil"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NFTClass holds the schema definition for the NFTClass entity.
type NFTClass struct {
	ent.Schema
}

// Fields of the BookNFT.
func (NFTClass) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").NotEmpty().Immutable(),
		field.Uint64("staked_amount").GoType(typeutil.Typ).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("staked_amount")...),
		field.Time("last_staked_at").Default(time.Now()),
		field.Uint64("number_of_stakers").Default(0),
	}
}

// Edges of the BookNFT.
func (NFTClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("accounts", Account.Type).
			Through("stakings", Staking.Type).
			Ref("nft_classes").
			Immutable(),
		edge.From("staking_events", StakingEvent.Type).
			Ref("nft_class").
			Immutable(),
	}
}

func (NFTClass) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address").Unique(),
	}
}
