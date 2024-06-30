package main

import "context"

type service struct {
	Store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{}
}

// implementing interface
func (s *service) CreateOrder(context.Context) error {
	return nil //does nothing for now
}
