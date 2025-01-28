package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NFT holds the schema definition for the NFT entity.
type NFT struct {
	ent.Schema
}

func (NFT) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "nfts"},
	}
}

// Fields of the NFT.
func (NFT) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("token_id").Unique(),
		field.String("token_url").Optional(),
		// Storing the content of the token_url as raw json
		field.JSON("raw", map[string]any{}).Optional(),
		// START Prepopulate field
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.JSON("image", map[string]any{}).Optional(),
		field.JSON("attributes", map[string]any{}).Optional(),
		// END Prepopulate field
		field.String("owner_address").NotEmpty(),
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

func (NFT) Index() []ent.Index {
	return []ent.Index{
		index.Fields("owner_address"),
	}
}
