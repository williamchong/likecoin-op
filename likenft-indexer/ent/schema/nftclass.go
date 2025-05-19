package schema

import (
	"math/big"
	"slices"

	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/evm/model"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NFTClass holds the schema definition for the NFTClass entity.
type NFTClass struct {
	ent.Schema
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
