package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "time"
)

func handleError(err error) {
    log.Println(err.Error())
    msg, _ := json.Marshal(err.Error())
    http.Error(writer, string(msg), 500)
}

func validatePosParams(yKey, xKey string) (int, int, bool) {
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
    var aiMove = Pos{-1, -1, 0}
    var line *[]Coord
    var notification string
    timerStart := time.Now().UnixNano()
    writer = w
    isWin := false

    query := r.URL.Query()
    y, x, valid := validatePosParams(query.Get("y"), query.Get("x"))
    if valid == false {
        return
    }

    board[y][x] = HUMAN
    if doubleThreeRule && freeThreeHuman && checkForFreeThree(y, x, HUMAN) {
        board[y][x] = EMPTY
        notification = fmt.Sprintf("No double-three")
    } else {
        isWin, line = fiveInARow(y, x, HUMAN)
        if isWin == false {
            aiMove = minimax(AI, Depth)
            board[aiMove.Y][aiMove.X] = AI
            if doubleThreeRule {
                resetFreeThrees()
            }
            isWin, line = fiveInARow(aiMove.Y, aiMove.X, AI)
        } else {
            notification = "Blue won"
        }
        if isWin == true {
            win = line
            if notification == "" {
                notification = "Red won"
            }
        }

        timerEnd := time.Now().UnixNano()
        notification = fmt.Sprintf("Time: %f %s",
            float64(timerEnd-timerStart)/1e9, notification)
    }

    data, _ := json.Marshal(struct {
        Y            int      `json:"y"`
        X            int      `json:"x"`
        Win          *[]Coord `json:"win"`
        Notification string   `json:"notification"`
    }{aiMove.Y, aiMove.X, win, notification})
    fmt.Fprintln(w, string(data))
}

func apiBoard(w http.ResponseWriter, r *http.Request) {
    data, _ := json.Marshal(struct {
        Board       *[BoardHeight][BoardWidth]int `json:"board"`
        Win         *[]Coord                      `json:"win"`
        Depth       int                           `json:"depth"`
        Moves       int                           `json:"moves"`
        DoubleThree bool                          `json:"double_three"`
    }{&board, win, Depth, MovesCheck, doubleThreeRule})
    fmt.Fprintln(w, string(data))
}

func apiReset(w http.ResponseWriter, r *http.Request) {
    board = [BoardHeight][BoardWidth]int{}
    board[BoardHeight/2][BoardWidth/2] = AI
    win = nil

    if doubleThreeRule {
        freeThreeAI = false
        freeThreeHuman = false
    }

    data, _ := json.Marshal(board)
    fmt.Fprintln(w, string(data))
}

func validateDifficultyParams(depth, moves string) (int, int, bool) {
    if depth == "" || moves == "" {
        handleError(errors.New("no get params"))
        return 0, 0, false
    }
    newDepth, errDepth := strconv.Atoi(depth)
    newMoves, errMoves := strconv.Atoi(moves)
    if errDepth != nil || errMoves != nil ||
        newDepth < 1 || newMoves < 1 ||
        newDepth > maxDepth || newMoves > maxMovesCheck {
        if errDepth != nil {
            handleError(errDepth)
        } else if errMoves != nil {
            handleError(errMoves)
        } else {
            handleError(errors.New("invalid params"))
        }
        return 0, 0, false
    }
    return newDepth, newMoves, true
}

func apiDifficulty(w http.ResponseWriter, r *http.Request) {
    writer = w

    depth := r.URL.Query().Get("depth")
    moves := r.URL.Query().Get("moves")
    newDepth, newMoves, valid := validateDifficultyParams(depth, moves)
    if valid == false {
        return
    }

    Depth = newDepth
    MovesCheck = newMoves
}

func apiDoubleThree(w http.ResponseWriter, r *http.Request) {
    writer = w

    check := r.URL.Query().Get("check")

    if check == "" {
        handleError(errors.New("no get params"))
        return
    }

    if check != "true" && check != "false" {
        handleError(errors.New("invalid params"))
        return
    }

    if check == "true" {
        doubleThreeRule = true
    } else {
        doubleThreeRule = false
    }
}

func scenario1(w http.ResponseWriter, r *http.Request) {
    board = [BoardHeight][BoardWidth]int{
        {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
        {0, 0, 0, 0, 2, 0, 0, 1, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        {0, 0, 0, 0, 1, 1, 0, 2, 2, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0},
        {0, 0, 1, 1, 2, 2, 1, 2, 1, 2, 2, 0, 0, 0, 1, 0, 0, 0, 0},
        {0, 1, 2, 2, 2, 1, 2, 1, 1, 2, 2, 1, 2, 2, 2, 0, 0, 0, 0},
        {0, 0, 2, 1, 0, 1, 2, 1, 1, 1, 2, 2, 1, 2, 2, 1, 2, 2, 0},
        {0, 1, 2, 2, 1, 2, 1, 0, 2, 2, 1, 1, 1, 1, 2, 0, 1, 0, 0},
        {2, 0, 1, 2, 2, 2, 1, 1, 1, 2, 2, 1, 1, 1, 2, 1, 0, 0, 0},
        {2, 0, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 2, 1, 1, 2, 1, 0, 0},
        {2, 1, 1, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 1, 1, 2, 2, 0, 0},
        {0, 0, 1, 1, 2, 1, 1, 0, 1, 2, 2, 1, 2, 2, 2, 2, 1, 1, 0},
        {0, 1, 2, 1, 2, 2, 1, 1, 2, 1, 1, 1, 1, 2, 1, 2, 2, 0, 0},
        {2, 0, 1, 1, 1, 2, 2, 1, 1, 2, 2, 2, 1, 2, 1, 1, 2, 1, 1},
        {0, 0, 0, 2, 0, 2, 1, 1, 2, 2, 1, 0, 2, 1, 0, 0, 1, 2, 0},
        {0, 0, 0, 2, 2, 1, 1, 2, 1, 0, 2, 1, 0, 2, 1, 0, 2, 2, 1},
        {0, 0, 0, 2, 1, 2, 1, 2, 0, 0, 0, 0, 1, 2, 2, 2, 2, 1, 0},
        {0, 0, 1, 2, 1, 2, 2, 2, 2, 1, 0, 1, 2, 2, 1, 1, 2, 0, 0},
        {0, 0, 0, 1, 0, 2, 0, 1, 0, 0, 0, 2, 2, 1, 0, 0, 1, 0, 0},
        {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
    }
    win = nil
    data, _ := json.Marshal(board)
    fmt.Fprintln(w, string(data))
}

func scenario2(w http.ResponseWriter, r *http.Request) {
    board = [BoardHeight][BoardWidth]int{
        {1, 0, 1, 2, 2, 1, 0, 1, 2, 1, 1, 2, 2, 1, 2, 1, 1, 1, 2},
        {2, 1, 2, 1, 2, 2, 1, 1, 1, 2, 2, 2, 1, 2, 2, 1, 2, 1, 2},
        {0, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 2, 1, 1, 2, 2, 2, 1},
        {1, 2, 1, 1, 2, 2, 1, 2, 1, 2, 2, 1, 2, 1, 1, 2, 1, 2, 0},
        {2, 1, 2, 2, 2, 1, 2, 1, 1, 2, 2, 1, 2, 2, 2, 1, 0, 1, 1},
        {1, 1, 2, 1, 0, 1, 2, 1, 1, 1, 2, 2, 1, 2, 2, 1, 2, 2, 1},
        {1, 1, 2, 2, 1, 2, 1, 1, 2, 2, 1, 1, 1, 1, 2, 2, 1, 2, 2},
        {2, 2, 1, 2, 2, 2, 1, 1, 1, 2, 2, 1, 1, 1, 2, 1, 1, 1, 1},
        {2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 2, 1, 1, 2, 1, 2, 1},
        {2, 1, 1, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 1, 1, 2, 2, 2, 1},
        {0, 1, 1, 1, 2, 1, 1, 0, 1, 2, 2, 1, 2, 2, 2, 2, 1, 1, 2},
        {1, 1, 2, 1, 2, 2, 1, 1, 2, 1, 1, 1, 1, 2, 1, 2, 2, 1, 1},
        {2, 2, 1, 1, 1, 2, 2, 1, 1, 2, 2, 2, 1, 2, 1, 1, 2, 1, 1},
        {2, 2, 2, 2, 1, 2, 1, 1, 2, 2, 1, 0, 2, 1, 0, 2, 1, 2, 2},
        {2, 1, 1, 2, 2, 1, 1, 2, 1, 2, 2, 1, 2, 2, 1, 1, 2, 2, 1},
        {2, 1, 2, 2, 1, 2, 1, 2, 1, 2, 1, 1, 1, 2, 2, 2, 2, 1, 1},
        {1, 1, 1, 2, 1, 2, 2, 2, 2, 1, 2, 1, 2, 2, 1, 1, 2, 2, 0},
        {2, 2, 2, 1, 0, 2, 2, 1, 2, 2, 1, 2, 2, 1, 2, 0, 1, 1, 1},
        {2, 1, 1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2, 1, 2, 2, 0, 2},
    }
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
    http.HandleFunc("/api/difficulty", apiDifficulty)
    http.HandleFunc("/api/rules/double_three", apiDoubleThree)
    http.HandleFunc("/api/scenario1", scenario1)
    http.HandleFunc("/api/scenario2", scenario2)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
