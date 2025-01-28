package schema

import (
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
		field.String("cosmos_address").Nillable().Optional(),
		field.String("evm_address").Unique(),
		field.String("likeid").Unique().Nillable().Optional(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nft_classes", NFTClass.Type),
		edge.To("nfts", NFT.Type),
	}
}

func (Account) Index() []ent.Index {
	return []ent.Index{
		index.Fields("cosmos_address"),
	}
}
