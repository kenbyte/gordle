package main

import "github.com/google/uuid"

type Network struct {
    ping   uint
    online bool
}
type Player struct {
    ID           uuid.UUID `json:"id"`
    Name         string    `json:"name"`
    GameHistory  []string  `json:"game_history"`
    CreationDate string    `json:"creation_date"`
    Score        uint      `json:"score"`
    Rank         uint      `json:"rank"`
}

func NewPlayer(name string) *Player {
    return &Player{
        ID:          uuid.New(),
        Name:        name,
        GameHistory: []string{},
        Score:       0,
    }
}
