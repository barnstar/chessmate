package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type ChessMateServer struct {
	gm GameManager
}

func (s *ChessMateServer) RunServer() {
	http.HandleFunc("/", s.handleClient)
	http.HandleFunc("/create", s.handleCreateGame)
	http.HandleFunc("/end", s.handleEndGame)
	http.HandleFunc("/join", s.handleJoinGame)

	fmt.Println("Server starting on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func (s *ChessMateServer) handleClient(w http.ResponseWriter, r *http.Request) {
	// Specify the path to the index.html file
	filePath := "../client/index.html"
	// Serve the file
	http.ServeFile(w, r, filePath)
}

// Creates a new game and returns its ID
func (s *ChessMateServer) handleCreateGame(w http.ResponseWriter, r *http.Request) {
	id := s.gm.CreateGame()
	fmt.Fprintf(w, "%s", id)
	// TODO: Respond back with the game ID. This should be a JSON object
}

// Creates a new game and returns its ID
func (s *ChessMateServer) handleEndGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the endGame API!")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin function to allow connections from any origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (s *ChessMateServer) handleJoinGame(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	gameId := gameId(params["gameId"][0]) // gameId=<abcd>
	color := color(params["color"][0])    // color=white
	name := params["name"][0]             // color=white

	fmt.Printf("Request to join game as %v with ID: %v", color, gameId)

	if gameId == "" || color == "" || !color.isValid() {
		fmt.Println("Invalid request")
		http.Error(w, "Invalid request parameters", http.StatusBadRequest)
		return
	}

	if name == "" {
		name = "Player"
	}

	if _, exists := s.gm.games[gameId]; !exists {
		fmt.Println("Game does not exist")
		http.Error(w, "Game does not exist", http.StatusNotFound)
		return
	}

	// Upgrade the HTTP server connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		http.Error(w, "Unable to start websocket", http.StatusInternalServerError)
		return
	}

	_, err = s.gm.JoinGame(gameId, conn, color, name)
	if err != nil {
		fmt.Println("Error joining game:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// After upgrading to WebSocket, you can use conn to read and write messages
	fmt.Fprintln(w, "WebSocket connection established")

	go func() {
		s.listenForEvents(conn, gameId)
	}()
}

func (s *ChessMateServer) listenForEvents(conn *websocket.Conn, id gameId) {
	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) ||
				err == websocket.ErrCloseSent {
				fmt.Println("WebSocket connection closed")
				break
			}
			// TODO: Tell the other party that the other player has left the game?
			//       Maybe just a genereic "update gamestate for everybody"
		}
		fmt.Printf("Message received: %s\n", message)
		s.gm.HandleMessage(id, message, msgType)
	}
}
