package main

import "context"

type store struct {
	//db instance
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	return nil //for now
}
