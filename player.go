package main

import "github.com/google/uuid"

type Player struct {
	ID          uuid.UUID
	Name        string
	GameHistory []string
}

func NewPlayer(name string) *Player {
	return &Player{
		ID:          uuid.New(),
		Name:        name,
		GameHistory: []string{},
	}
}
