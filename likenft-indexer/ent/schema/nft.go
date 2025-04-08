package schema

import (
	"slices"

	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/evm/model"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

//go:generate go generate ../../ent

// NFT holds the schema definition for the NFT entity.
type NFT struct {
	ent.Schema
}

func (NFT) Annotations() []schema.Annotation {
	return slices.Concat(
		[]schema.Annotation{
			entsql.Annotation{Table: "nfts"},
		},
		typeutil.Uint64Annotations("token_id"),
	)
}

// Fields of the NFT.
func (NFT) Fields() []ent.Field {
	return []ent.Field{
		field.String("contract_address").MaxLen(42).NotEmpty(),
		field.Uint64("token_id").
			GoType(typeutil.Uint64(0)).
			SchemaType(typeutil.Uint64SchemaType).
			ValueScanner(typeutil.Uint64ValueScanner),
		field.String("token_uri").Nillable().Optional(),
		// START Prepopulate field
		field.String("image").Nillable().Optional(),
		field.String("image_data").Nillable().Optional(),
		field.String("external_url").Nillable().Optional(),
		field.String("description").Nillable().Optional(),
		field.String("name").Nillable().Optional(),
		field.JSON("attributes", []model.ERC721MetadataAttribute{}).Optional(),
		field.String("background_color").Nillable().Optional(),
		field.String("animation_url").Nillable().Optional(),
		field.String("youtube_url").Nillable().Optional(),
		// END Prepopulate field
		field.String("owner_address").MaxLen(42).NotEmpty(),
		field.Time("minted_at"),
		field.Time("updated_at"),
	}
}

// Edges of the NFT.
func (NFT) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Account.Type).
			Ref("nfts").
			Unique(),
		edge.From("class", NFTClass.Type).
			Ref("nfts").
			Unique(),
	}
}

func (NFT) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("contract_address", "token_id").Unique(),
		index.Fields("owner_address"),
	}
}
