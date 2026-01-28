package tictactoe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type GameState struct {
	Board    [9]string `json:"board"`
	GameOver bool      `json:"gameOver"`
	Winner   string    `json:"winner"`
	Message  string    `json:"message"`
}

var tictactoeTemplate *template.Template

func init() {
	var err error
	tictactoeTemplate, err = template.ParseFiles("templates/tictactoe.html")
	if err != nil {
		fmt.Printf("Error parsing tictactoe template: %v\n", err)
	}
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/tictactoe", handleTicTacToePage)
	mux.HandleFunc("/tictactoe/api/move", handleMove)
	mux.HandleFunc("/tictactoe/api/new-game", handleNewGame)
}

func handleTicTacToePage(w http.ResponseWriter, r *http.Request) {
	initialGame := GameState{
		Board:    [9]string{},
		GameOver: false,
		Winner:   "",
		Message:  "Your turn (X)",
	}
	for i := range initialGame.Board {
		initialGame.Board[i] = ""
	}
	tictactoeTemplate.Execute(w, initialGame)
}

func handleMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var gameState GameState
	err := json.NewDecoder(r.Body).Decode(&gameState)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check for winner or draw
	gameState.Winner = checkWinner(gameState.Board[:])
	if gameState.Winner != "" {
		gameState.GameOver = true
		gameState.Message = fmt.Sprintf("%s wins!", gameState.Winner)
		json.NewEncoder(w).Encode(gameState)
		return
	}

	if isBoardFull(gameState.Board[:]) {
		gameState.GameOver = true
		gameState.Message = "It's a draw!"
		json.NewEncoder(w).Encode(gameState)
		return
	}

	// Computer move
	computerMove(&gameState)

	gameState.Winner = checkWinner(gameState.Board[:])
	if gameState.Winner != "" {
		gameState.GameOver = true
		gameState.Message = fmt.Sprintf("%s wins!", gameState.Winner)
		json.NewEncoder(w).Encode(gameState)
		return
	}

	if isBoardFull(gameState.Board[:]) {
		gameState.GameOver = true
		gameState.Message = "It's a draw!"
		json.NewEncoder(w).Encode(gameState)
		return
	}

	gameState.Message = "Your turn (X)"
	json.NewEncoder(w).Encode(gameState)
}

func handleNewGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newGame := GameState{
		Board:    [9]string{},
		GameOver: false,
		Winner:   "",
		Message:  "Your turn (X)",
	}
	for i := range newGame.Board {
		newGame.Board[i] = ""
	}
	json.NewEncoder(w).Encode(newGame)
}

func checkWinner(board []string) string {
	// Check rows
	for i := 0; i < 9; i += 3 {
		if board[i] != "" && board[i] == board[i+1] && board[i+1] == board[i+2] {
			return board[i]
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if board[i] != "" && board[i] == board[i+3] && board[i+3] == board[i+6] {
			return board[i]
		}
	}

	// Check diagonals
	if board[0] != "" && board[0] == board[4] && board[4] == board[8] {
		return board[0]
	}
	if board[2] != "" && board[2] == board[4] && board[4] == board[6] {
		return board[2]
	}

	return ""
}

func isBoardFull(board []string) bool {
	for _, cell := range board {
		if cell == "" {
			return false
		}
	}
	return true
}

func computerMove(gameState *GameState) {
	// Simple AI: Try to win, then block player, then take center, then take corner, then take any
	emptySpots := getEmptySpots(gameState.Board[:])

	// Try to win
	for _, spot := range emptySpots {
		gameState.Board[spot] = "O"
		if checkWinner(gameState.Board[:]) == "O" {
			return
		}
		gameState.Board[spot] = ""
	}

	// Try to block
	for _, spot := range emptySpots {
		gameState.Board[spot] = "X"
		if checkWinner(gameState.Board[:]) == "X" {
			gameState.Board[spot] = "O"
			return
		}
		gameState.Board[spot] = ""
	}

	// Take center if available
	if gameState.Board[4] == "" {
		gameState.Board[4] = "O"
		return
	}

	// Take a corner
	corners := []int{0, 2, 6, 8}
	for _, corner := range corners {
		if gameState.Board[corner] == "" {
			gameState.Board[corner] = "O"
			return
		}
	}

	// Take any available
	if len(emptySpots) > 0 {
		gameState.Board[emptySpots[0]] = "O"
	}
}

func getEmptySpots(board []string) []int {
	var empty []int
	for i, cell := range board {
		if cell == "" {
			empty = append(empty, i)
		}
	}
	return empty
}
