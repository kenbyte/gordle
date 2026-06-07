package main

import (
	"time"

	"github.com/google/uuid"
)

type Lobby struct {
	Host       *Player
	Invite     *Player
	ReadyState map[uuid.UUID]bool
}

type Participant struct {
	Player   *Player
	Attempts []string
	Guesses  uint
	Solved   bool
}

type Game struct {
	ID         uuid.UUID `json:"id"`
	StartTime  time.Time `json:"start_time"`
	IsFinished bool      `json:"is_finished"`
	IsSingle   bool      `json:"is_single"`
	Word       string    `json:"word"`

	Participants []*Participant `json:"-"`
	Lobby        *Lobby         `json:"-"`
}

func NewSingleGame(isSingle bool, word string) *Game {
	return &Game{
		ID:           uuid.New(),
		StartTime:    time.Now(),
		IsFinished:   false,
		IsSingle:     isSingle,
		Word:         word,
		Participants: []*Participant{},
		Lobby:        &Lobby{},
	}
}
