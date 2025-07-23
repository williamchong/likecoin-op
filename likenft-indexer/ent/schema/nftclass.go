package schema

import (
	"math/big"
	"slices"

	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/evm/model"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// NFTClass holds the schema definition for the NFTClass entity.
type NFTClass struct {
	ent.Schema
}

func (NFTClass) Mixin() []ent.Mixin {
	return []ent.Mixin{
		nftClassAcquireBookNFTEventsMixin{},
	}
}

// Fields of the NFTClass.
func (NFTClass) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Unique(),
		field.String("name").NotEmpty(),
		field.String("symbol").NotEmpty(),
		field.String("owner_address").Nillable().Optional(),
		// Minter addresses is commonly bookstore addresses
		field.JSON("minter_addresses", []string{}).Optional(),
		field.Uint64("total_supply").GoType(&big.Int{}).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(field.TextValueScanner[*big.Int]{}),
		field.Uint64("max_supply").GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
		// Raw metadata from the contract
		field.JSON("metadata", &model.ContractLevelMetadata{}).Optional(),
		// Start Prepopulate fields
		field.String("banner_image"),
		field.String("featured_image"),
		// End Prepopulate fields
		field.String("deployer_address").NotEmpty(),
		field.Uint64("deployed_block_number").GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
		field.Uint64("latest_event_block_number").GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
		field.Bool("disabled_for_indexing").Default(false),
		field.String("disabled_for_indexing_reason").Optional(),
		field.Time("minted_at"),
		field.Time("updated_at"),
	}
}

// Edges of the NFTClass.
func (NFTClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nfts", NFT.Type),
		edge.From("owner", Account.Type).
			Ref("nft_classes").
			Unique(),
	}
}

func (NFTClass) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_address"),
		index.Fields("deployer_address"),
	}
}

func (NFTClass) Annotations() []schema.Annotation {
	return slices.Concat(
		typeutil.Uint64Annotations("total_supply"),
		typeutil.Uint64Annotations("max_supply"),
		typeutil.Uint64Annotations("deployed_block_number"),
		typeutil.Uint64Annotations("latest_event_block_number"),
	)
}

// --- mixins

type nftClassAcquireBookNFTEventsMixin struct {
	mixin.Schema
}

func (nftClassAcquireBookNFTEventsMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Float("acquire_book_nft_events_weight").Default(1).Min(0),
		field.Time("acquire_book_nft_events_last_processed_time").Optional().Nillable(),
		field.Float("acquire_book_nft_events_eta").SchemaType(map[string]string{
			dialect.Postgres: "numeric",
		}).Optional().Nillable(),
		field.Enum("acquire_book_nft_events_status").Values(
			"enqueueing",
			"enqueued",
			"enqueue_failed",
			"processing",
			"completed",
			"failed",
		).Optional().Nillable(),
		field.String("acquire_book_nft_events_failed_reason").Optional().Nillable(),
		field.Int("acquire_book_nft_events_failed_count").Default(0),
	}
}

func (nftClassAcquireBookNFTEventsMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("acquire_book_nft_events_eta"),
		index.Fields("acquire_book_nft_events_status"),
	}
}
