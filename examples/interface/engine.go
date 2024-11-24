package main

import "github.com/nanostack-dev/do"

type Engine interface{}

type engineImplem struct {
}

func NewEngine(i do.Injector) (*engineImplem, error) {
	return &engineImplem{}, nil
}
