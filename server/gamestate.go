package server

type color string
type gameID string

type piece struct {
	color        // black or white
	name  string // pawn, rook, knight, bishop, queen, king
}

type GameState struct {
	board [8][8]piece
}

type GameManager struct {
	games map[gameID]GameState
}
