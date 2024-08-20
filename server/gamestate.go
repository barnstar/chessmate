package server

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type color string
type gameId string

func (c *color) isValid() bool {
	return *c == "white" || *c == "black"
}

type Piece struct {
	Color color  `json:"color"`
	Name  string `json:"name"`
}
type player struct {
	Name   string          `json:"name"`
	Rating int             `json:"rating"`
	Conn   *websocket.Conn `json:"-"`
}

type GameState struct {
	Id          gameId      `json:"id"`
	Board       [8][8]Piece `json:"board"`
	WhitePlayer player      `json:"whitePlayer"`
	BlackPlayer player      `json:"blackPlayer"`
}

func (gs *GameState) SetPiece(x int, y int, p Piece) {
	gs.Board[x][y] = p
}

func (gs *GameState) serialize() (string, error) {
	jsonData, err := json.Marshal(gs)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
