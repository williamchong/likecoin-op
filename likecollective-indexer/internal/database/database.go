package database

import "likecollective-indexer/ent"

type Service interface {
	Client() *ent.Client
}

type service struct {
	client *ent.Client
}

func New() Service {
	return &service{
		client: ent.NewClient(),
	}
}

func (s *service) Client() *ent.Client {
	return s.client
}
