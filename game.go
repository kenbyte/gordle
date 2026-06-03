package main

import (
	"time"

	"github.com/google/uuid"
)


type Game struct {
	ID         uuid.UUID
	Timer      time.Time
	isFinished bool
	isSingle   bool
	word       string
	player     []Player
}

func NewGame(isSingle bool, word string) *Game {
	return &Game{
		ID:         uuid.New(),
		Timer:      time.Now(),
		isFinished: false,
		isSingle:   isSingle,
		word:       word,
	}
}

func (g *Game) CloseGame() {
	g.isFinished = true
}
