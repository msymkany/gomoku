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
    var aiMove = Move{nil, 0}
    var isWin = false
    var debug *Debug = nil
    if debugMode {
        debug = &Debug{0, 0, nil, 0, -1, []*Debug{}}
    }
    var line *[]Coord
    var notification string

    timerStart := time.Now().UnixNano()
    writer = w

    query := r.URL.Query()
    y, x, valid := validatePosParams(query.Get("y"), query.Get("x"))
    if valid == false {
        return
    }

    pos := &Position{captureMove(y, x, HUMAN), y, x}
    if doubleThree(pos, HUMAN) {
        notification = fmt.Sprintf("No double-three")
    } else {
        makeMove(pos, HUMAN)
        if captures[HUMAN] == 10 {
            notification = "Blue won"
            winByCapture = 1
        } else {
            isWin, line = fiveInARow(y, x, HUMAN)
            if isWin == false {
                aiMove = minimax(AI, Depth, debug)
                makeMove(aiMove.pos, AI)
                if doubleThreeRule {
                    resetFreeThrees()
                }
                if captures[AI] == 10 {
                    notification = "Red won"
                    winByCapture = 2
                } else {
                    isWin, line = fiveInARow(aiMove.pos.Y, aiMove.pos.X, AI)
                    if isWin == true {
                        notification = "Red won"
                    }
                }
            } else {
                notification = "Blue won"
            }
        }

        if isWin == true {
            win = line
        }

        timerEnd := time.Now().UnixNano()
        notification = fmt.Sprintf("Time: %f %s",
            float64(timerEnd-timerStart)/1e9, notification)
    }

    data, _ := json.Marshal(struct {
        BluePos      *Position   `json:"blue_pos"`
        RedPos       *Position   `json:"red_pos"`
        Win          *[]Coord    `json:"win"`
        WinByCapture int         `json:"win_by_capture"`
        Notification string      `json:"notification"`
        Captures     map[int]int `json:"captures"`
        Debug        *Debug      `json:"debug"`
    }{
        pos,
        aiMove.pos,
        win,
        winByCapture,
        notification,
        captures,
        debug,
    })
    fmt.Fprintln(w, string(data))
}

func apiBoard(w http.ResponseWriter, r *http.Request) {
    data, _ := json.Marshal(struct {
        Board        *[BoardHeight][BoardWidth]int `json:"board"`
        Win          *[]Coord                      `json:"win"`
        WinByCapture int                           `json:"win_by_capture"`
        Depth        int                           `json:"depth"`
        Moves        int                           `json:"moves"`
        DoubleThree  bool                          `json:"double_three"`
        CaptureRule  bool                          `json:"capture_rule"`
        Captures     map[int]int                   `json:"captures"`
        DebugMode    bool                          `json:"debug_mode"`
    }{
        &board,
        win,
        winByCapture,
        Depth,
        MovesCheck,
        doubleThreeRule,
        captureRule,
        captures,
        debugMode,
    })
    fmt.Fprintln(w, string(data))
}

func apiReset(w http.ResponseWriter, r *http.Request) {
    board = [BoardHeight][BoardWidth]int{}
    board[BoardHeight/2][BoardWidth/2] = AI
    win = nil
    winByCapture = 0

    if doubleThreeRule {
        freeThreeAI = false
        freeThreeHuman = false
    }

    captures[AI] = 0
    captures[HUMAN] = 0

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
    if valid {
        Depth = newDepth
        MovesCheck = newMoves
    }
}

func changeSetting(w http.ResponseWriter, r *http.Request, boolVar *bool) {
    writer = w

    check := r.URL.Query().Get("check")

    if check != "" && (check == "true" || check == "false") {
        if check == "true" {
            *boolVar = true
        } else {
            *boolVar = false
        }
    } else if check == "" {
        handleError(errors.New("no get params"))
    } else {
        handleError(errors.New("invalid params"))
    }
}

func apiDoubleThree(w http.ResponseWriter, r *http.Request) {
    changeSetting(w, r, &doubleThreeRule)
}

func apiCapture(w http.ResponseWriter, r *http.Request) {
    changeSetting(w, r, &captureRule)
}

func apiDebug(w http.ResponseWriter, r *http.Request) {
    changeSetting(w, r, &debugMode)
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
    http.HandleFunc("/api/settings/difficulty", apiDifficulty)
    http.HandleFunc("/api/settings/double_three", apiDoubleThree)
    http.HandleFunc("/api/settings/capture", apiCapture)
    http.HandleFunc("/api/settings/debug", apiDebug)
    http.HandleFunc("/api/scenario1", scenario1)
    http.HandleFunc("/api/scenario2", scenario2)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
