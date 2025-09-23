package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// StakingEventsHyperTable holds the schema definition for the StakingEventsHyperTable entity.
type StakingEventsHyperTable struct {
	ent.Schema
}

func (StakingEventsHyperTable) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("staking_events_hyper_table"),
		entsql.Skip(),
	}
}

// Fields of the StakingEventsHyperTable.
func (StakingEventsHyperTable) Fields() []ent.Field {
	return StakingEvent{}.Fields()
}
