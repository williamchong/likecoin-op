package schema

import (
	"likecollective-indexer/ent/schema/typeutil"
	"time"

	"entgo.io/ent"
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
		field.String("transaction_hash").
			NotEmpty().
			Immutable(),
		field.Uint("transaction_index").
			Immutable(),
		field.Uint64("block_number").
			GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner).
			Annotations(typeutil.Uint64Annotations("block_number")...).
			Immutable(),
		field.Uint("log_index").
			Immutable(),
		field.Enum("event_type").
			Values(
				"staked",
				"unstaked",
				"reward_claimed",
				"reward_deposited",
				"reward_deposit_distributed",
				"all_rewards_claimed",
			).
			Default("staked").
			Immutable(),
		field.String("nft_class_address").
			NotEmpty().
			Immutable(),
		field.String("account_evm_address").
			NotEmpty().
			Immutable(),
		field.Uint64("staked_amount_added").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("staked_amount_added")...).
			Immutable(),
		field.Uint64("staked_amount_removed").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("staked_amount_removed")...).
			Immutable(),
		field.Uint64("pending_reward_amount_added").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("pending_reward_amount_added")...).
			Immutable(),
		field.Uint64("pending_reward_amount_removed").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("pending_reward_amount_removed")...).
			Immutable(),
		field.Uint64("claimed_reward_amount_added").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("claimed_reward_amount_added")...).
			Immutable(),
		field.Uint64("claimed_reward_amount_removed").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("claimed_reward_amount_removed")...).
			Immutable(),
		field.Time("datetime").Default(time.Now()).
			Immutable(),
	}
}

// Edges of the StakingEvent.
func (StakingEvent) Edges() []ent.Edge {
	return nil
}

func (StakingEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_type"),
	}
}
