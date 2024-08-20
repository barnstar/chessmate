package server

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type GameManager struct {
	games map[gameId]GameState
}

func (gm *GameManager) CreateGame() gameId {
	var gs GameState
	id := gameId(uuid.New().String())
	gm.games[id] = gs
	fmt.Printf("Created game with ID %s\n", id)
	return id
}

func (gm *GameManager) EndGame(id gameId) {
	if gs, exists := gm.games[id]; !exists {
		fmt.Printf("Game with ID %s does not exist\n", id)
		return
	} else {
		fmt.Printf("Ending game with ID %s\n", id)
		if gs.WhitePlayer.Conn != nil {
			gs.WhitePlayer.Conn.Close()
		}
		if gs.BlackPlayer.Conn != nil {
			gs.BlackPlayer.Conn.Close()
		}
	}
	delete(gm.games, id)
}

func (gm *GameManager) JoinGame(id gameId, conn *websocket.Conn, color color, name string) (*GameState, error) {
	if gs, exists := gm.games[id]; !exists {
		fmt.Printf("Game with ID %s does not exist\n", id)
		return nil, fmt.Errorf("%s does not exist", id)
	} else {
		fmt.Printf("Joining game with ID %s\n", id)
		if color == "white" && gs.WhitePlayer.Conn == nil {
			gs.WhitePlayer.Conn = conn
			gs.WhitePlayer.Name = name
			fmt.Printf("Joining game with ID %s as %s playing %s\n", id, name, color)
		} else if color == "black" && gs.BlackPlayer.Conn == nil {
			gs.BlackPlayer.Conn = conn
			gs.BlackPlayer.Name = name
			fmt.Printf("Joining game with ID %s as %s playing %s\n", id, name, color)
		} else {
			fmt.Printf("Game with ID %s is full\n", id)
			return nil, fmt.Errorf("%s is full", id)
		}
		return &gs, nil
	}
}

func (gm *GameManager) HandleMessage(id gameId, message []byte, msgType int) {
	if _, exists := gm.games[id]; !exists {
		fmt.Printf("Game with ID %s does not exist\n", id)
		return
	} else {
		// TODO: Determine the message object type from msgType
		// TODO: Deserialize the bytes into the proper message object
		// TODO: Validate the message
		// TODO: Update the gamestate
		// TODO: Send gamestate updated message to all connected peers
		fmt.Printf("Message received for game with ID %s\n", id)
		fmt.Printf("Message: %s\n", message)
	}
}

func (gm *GameManager) GetGames() map[gameId]GameState {
	return gm.games
}
