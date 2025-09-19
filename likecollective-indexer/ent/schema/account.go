package schema

import (
	"likecollective-indexer/ent/schema/typeutil"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("evm_address").NotEmpty().Immutable(),
		field.Uint64("staked_amount").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("staked_amount")...),
		field.Uint64("pending_reward_amount").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("pending_reward_amount")...),
		field.Uint64("claimed_reward_amount").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("claimed_reward_amount")...),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nft_classes", NFTClass.Type).
			Through("stakings", Staking.Type).
			Immutable(),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("evm_address").Unique(),
	}
}
