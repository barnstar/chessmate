package server

import (
	"fmt"
	"net/http"
)

type ChessMateServer struct {
	GameManager GameManager
}

func (s *ChessMateServer) RunServer() {
	http.HandleFunc("/", s.handleClient)
	http.HandleFunc("/move", s.handleMove)
	http.HandleFunc("/getstate", s.handleGetstate)
	http.HandleFunc("/setstate", s.handleSetstate)
	http.HandleFunc("/create", s.handleCreateGame)
	http.HandleFunc("/end", s.handleEndGame)

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

// Handles a move request for a given game
func (s *ChessMateServer) handleMove(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the move API!")
}

// Returns the current state of a given game
func (s *ChessMateServer) handleGetstate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the getState API!")
}

// Sets the state of a given game
func (s *ChessMateServer) handleSetstate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the setState API!")
}

// Creates a new game and returns its ID
func (s *ChessMateServer) handleCreateGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the create API!")
}

// Creates a new game and returns its ID
func (s *ChessMateServer) handleEndGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the endGame API!")
}
