package schema

import (
	"likecollective-indexer/ent/schema/typeutil"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// StakingEvent holds the schema definition for the StakingEvent entity.
type StakingEvent struct {
	ent.Schema
}

// Fields of the StakingEvent.
func (StakingEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("event_type").
			Values(
				"staked",
				"unstaked",
				"reward_added",
				"reward_claimed",
				"reward_deposited",
				"all_rewards_claimed",
			).
			Default("staked").
			Immutable(),
		field.Int("nft_class_id").
			Immutable(),
		field.Int("account_id").
			Immutable(),
		field.Uint64("staked_amount_added").GoType(typeutil.Typ).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("staked_amount_added")...).
			Immutable(),
		field.Uint64("staked_amount_removed").GoType(typeutil.Typ).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("staked_amount_removed")...).
			Immutable(),
		field.Uint64("reward_amount_added").GoType(typeutil.Typ).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("reward_amount_added")...).
			Immutable(),
		field.Uint64("reward_amount_removed").GoType(typeutil.Typ).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("reward_amount_removed")...).
			Immutable(),
		field.Time("datetime").Default(time.Now()).
			Immutable(),
	}
}

// Edges of the StakingEvent.
func (StakingEvent) Edges() []ent.Edge {
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

func (StakingEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_type"),
	}
}
