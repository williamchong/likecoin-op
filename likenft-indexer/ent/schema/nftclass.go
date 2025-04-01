package schema

import (
	"entgo.io/ent"
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
		field.Int("total_supply").NonNegative(),
		// Raw metadata from the contract
		field.JSON("metadata", map[string]any{}).Optional(),
		// Start Prepopulate fields
		field.String("banner_image"),
		field.String("featured_image"),
		// End Prepopulate fields
		field.String("deployer_address").NotEmpty(),
		field.String("deployed_block_number").NotEmpty(),
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
