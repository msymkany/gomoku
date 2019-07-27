package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handleError(err error) {
	log.Println(err.Error())
	msg, _ := json.Marshal(err.Error())
	http.Error(writer, string(msg), 500)
}

func validGetParams(yKey, xKey string) (int, int, bool) {
	if yKey == "" || xKey == "" {
		handleError(errors.New("no get params"))
		return 0, 0, false
	}
	y, errY := strconv.Atoi(yKey)
	x, errX := strconv.Atoi(xKey)
	if errY != nil || errX != nil ||
		x < 0 || y < 0 || x > 18 || y > 18 ||
		board[y][x] != EMPTY {
		if errY != nil {
			handleError(errY)
		} else if errX != nil {
			handleError(errX)
		} else {
			handleError(errors.New("invalid params"))
		}
		return 0, 0, false
	}
	return y, x, true
}

func api(w http.ResponseWriter, r *http.Request) {
	writer = w
	var aiMove = Pos{-1, -1, 0}
	var line *[]Coord
	isWin := false
	query := r.URL.Query()
	y, x, valid := validGetParams(query.Get("y"), query.Get("x"))
	if valid == false {
		return
	}
	board[y][x] = HUMAN
	isWin, line = fiveInARow(y, x, HUMAN)
	if isWin == false {
		aiMove = minimax(AI, DEPTH)
		fmt.Println("main", aiMove)
		board[aiMove.Y][aiMove.X] = AI
		isWin, line = fiveInARow(aiMove.Y, aiMove.X, AI)
	}
	if isWin == true {
		win = line
	}
	data, _ := json.Marshal(struct {
		Y   int      `json:"y"`
		X   int      `json:"x"`
		Win *[]Coord `json:"win"`
	}{aiMove.Y, aiMove.X, win})
	fmt.Fprintln(w, string(data))
}

func apiBoard(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(struct {
		Board *[BoardHeight][BoardWidth]int `json:"board"`
		Win   *[]Coord                      `json:"win"`
	}{&board, win})
	fmt.Fprintln(w, string(data))
}

func apiReset(w http.ResponseWriter, r *http.Request) {
	board = [BoardHeight][BoardWidth]int{}
	board[BoardHeight/2][BoardWidth/2] = AI
	win = nil
	data, _ := json.Marshal(board)
	fmt.Fprintln(w, string(data))
}

func main() {
	board[BoardHeight/2][BoardWidth/2] = AI
	http.Handle("/", http.FileServer(http.Dir("front")))
	http.HandleFunc("/api", api)
	http.HandleFunc("/api/board", apiBoard)
	http.HandleFunc("/api/reset", apiReset)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
