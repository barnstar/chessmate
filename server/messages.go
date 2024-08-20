package server

// Sent in response to changes in the game state
type UpdateGameState struct {
	Id    string    `json:"messageId"`
	State GameState `json:"state"`
}

// Sent by clients to request a move be made
// from (fromX, fromY) to (toX, toY)
//
// If the move is valid, the gamestate will be updated
type Move struct {
	Id    string `json:"messageId"`
	FromX int    `json:"fromX"`
	FromY int    `json:"fromY"`
	ToX   int    `json:"toX"`
	ToY   int    `json:"toY"`
}

type MessageError struct {
	Id    string `json:"messageId"` // The message Id of the message that caused the error
	Error string `json:"error"`     // The error message
}

type errorMessages struct {
	InvalidMove string
}

var Errors = errorMessages{
	InvalidMove: "Invalid move",
}
