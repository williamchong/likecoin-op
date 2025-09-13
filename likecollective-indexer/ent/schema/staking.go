package schema

import (
	"likecollective-indexer/ent/schema/typeutil"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Staking holds the schema definition for the Staking entity.
type Staking struct {
	ent.Schema
}

// Fields of the Staking.
func (Staking) Fields() []ent.Field {
	return []ent.Field{
		field.Int("nft_class_id").Immutable(),
		field.Int("account_id").Immutable(),
		field.String("pool_share").NotEmpty(),
		field.Uint64("staked_amount").GoType(typeutil.Typ).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("staked_amount")...),
		field.Uint64("pending_reward_amount").GoType(typeutil.Typ).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("pending_reward_amount")...),
		field.Uint64("claimed_reward_amount").GoType(typeutil.Typ).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("claimed_reward_amount")...),
	}
}

// Edges of the Staking.
func (Staking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Field("account_id").
			Unique().
			Required().
			Immutable(),
		edge.To("nft_class", NFTClass.Type).
			Field("nft_class_id").
			Unique().
			Required().
			Immutable(),
	}
}

func (Staking) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("nft_class_id", "account_id").Unique(),
	}
}
