package schema

import (
	"likecollective-indexer/ent/schema/typeutil"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BookNFTDeltaTimeBucketMixin struct {
	ent.Schema
}

func (BookNFTDeltaTimeBucketMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Skip(),
	}
}

func (BookNFTDeltaTimeBucketMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Immutable(),
		field.String("evm_address").
			NotEmpty().
			Immutable(),
		field.Time("bucket").Immutable(),
		field.Uint64("staked_amount").
			GoType(typeutil.Uint256Type).
			SchemaType(typeutil.Uint256SchemaType).
			ValueScanner(typeutil.Uint256ValueScanner).
			Annotations(typeutil.Uint256Annotations("staked_amount")...).
			Immutable(),
		field.Time("last_staked_at").Immutable(),
		field.Uint64("number_of_stakers").Immutable(),
	}
}

type BookNFTDeltaTimeBucket7d struct {
	ent.Schema
}

func (BookNFTDeltaTimeBucket7d) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BookNFTDeltaTimeBucketMixin{},
	}
}

func (BookNFTDeltaTimeBucket7d) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("book_nft_delta_time_bucket_7d"),
	}
}

type BookNFTDeltaTimeBucket30d struct {
	ent.Schema
}

func (BookNFTDeltaTimeBucket30d) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BookNFTDeltaTimeBucketMixin{},
	}
}

func (BookNFTDeltaTimeBucket30d) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("book_nft_delta_time_bucket_30d"),
	}
}

type BookNFTDeltaTimeBucket1y struct {
	ent.Schema
}

func (BookNFTDeltaTimeBucket1y) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BookNFTDeltaTimeBucketMixin{},
	}
}

func (BookNFTDeltaTimeBucket1y) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("book_nft_delta_time_bucket_1y"),
	}
}
