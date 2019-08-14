package main

import (
    "container/heap"
)

func validX(x int) bool {
    if x >= 0 && x < BoardWidth {
        return true
    }
    return false
}

func validY(y int) bool {
    if y >= 0 && y < BoardHeight {
        return true
    }
    return false
}

func fiveInARow(y, x, player int) (bool, *[]Coord) {
    inARow := 0
    var line []Coord

    // Horizontal
    for tempX := x - 4; tempX <= x+4 && tempX < BoardWidth; tempX++ {
        if inARow >= 5 && validX(tempX) && board[y][tempX] != player {
            break
        }
        if validX(tempX) && board[y][tempX] == player {
            inARow++
            line = append(line, Coord{y, tempX})
        } else {
            inARow = 0
            line = []Coord{}
        }
    }
    if inARow >= 5 {
        return true, &line
    } else {
        inARow = 0
        line = []Coord{}
    }

    // Vertical
    for tempY := y - 4; tempY <= y+4 && tempY < BoardHeight; tempY++ {
        if inARow >= 5 && validY(tempY) && board[tempY][x] != player {
            break
        }
        if validY(tempY) && board[tempY][x] == player {
            inARow++
            line = append(line, Coord{tempY, x})
        } else {
            inARow = 0
            line = []Coord{}
        }
    }
    if inARow >= 5 {
        return true, &line
    } else {
        inARow = 0
        line = []Coord{}
    }

    // Diagonal 1
    for tempY, tempX := y-4, x-4; tempX <= x+4 && tempY < BoardHeight && tempX < BoardWidth; tempY, tempX = tempY+1, tempX+1 {
        if inARow >= 5 && validY(tempY) && validX(tempX) &&
            board[tempY][tempX] != player {
            break
        }
        if validY(tempY) && validX(tempX) && board[tempY][tempX] == player {
            inARow++
            line = append(line, Coord{tempY, tempX})
        } else {
            inARow = 0
            line = []Coord{}
        }
    }
    if inARow >= 5 {
        return true, &line
    } else {
        inARow = 0
        line = []Coord{}
    }

    // Diagonal 2
    for tempY, tempX := y+4, x-4; tempX <= x+4 && tempY >= 0 && tempX < BoardWidth; tempY, tempX = tempY-1, tempX+1 {
        if inARow >= 5 && validY(tempY) && validX(tempX) &&
            board[tempY][tempX] != player {
            break
        }
        if validY(tempY) && validX(tempX) && board[tempY][tempX] == player {
            inARow++
            line = append(line, Coord{tempY, tempX})
        } else {
            inARow = 0
            line = []Coord{}
        }
    }
    if inARow >= 5 {
        return true, &line
    } else {
        inARow = 0
        line = []Coord{}
    }

    return false, &line
}

func adjacentNotEmpty(y, x int) bool {
    if (y-1 >= 0 && x-1 >= 0 && board[y-1][x-1] != EMPTY) ||
        (y-1 >= 0 && board[y-1][x] != EMPTY) ||
        (y-1 >= 0 && x+1 < BoardWidth && board[y-1][x+1] != EMPTY) ||
        (x-1 >= 0 && board[y][x-1] != EMPTY) ||
        (x+1 < BoardWidth && board[y][x+1] != EMPTY) ||
        (y+1 < BoardHeight && x-1 >= 0 && board[y+1][x-1] != EMPTY) ||
        (y+1 < BoardHeight && board[y+1][x] != EMPTY) ||
        (y+1 < BoardHeight && x+1 < BoardWidth && board[y+1][x+1] != EMPTY) {
        return true
    }
    return false
}

func checkForAI(player int, scoreAI, scroreHuman float64) float64 {
    if player == AI {
        return scoreAI
    }
    return scroreHuman
}

func getScore(row, openEnds, player int) float64 {
    if row == 4 {
        if openEnds == 0 {
            return 0
        } else if openEnds == 1 {
            return checkForAI(player, 100000000, 50)
        } else if openEnds == 2 {
            return checkForAI(player, 100000000, 500000)
        }
    } else if row == 3 {
        if openEnds == 0 {
            return 0
        } else if openEnds == 1 {
            return checkForAI(player, 7, 5)
        } else if openEnds == 2 {
            return checkForAI(player, 10000, 50)
        }
    } else if row == 2 {
        if openEnds == 0 {
            return 0
        } else if openEnds == 1 {
            return 2
        } else if openEnds == 2 {
            return 5
        }
    } else if row == 1 {
        if openEnds == 0 {
            return 0
        } else if openEnds == 1 {
            return 0.5
        } else if openEnds == 2 {
            return 1
        }
    }
    return 200000000
}

func checkPlayer(score *float64, row, openEnds *int, y, x, player1, curPlayer int) {
    if board[y][x] == player1 {
        *row = *row + 1
    } else if board[y][x] == EMPTY && *row > 0 {
        *openEnds = *openEnds + 1
        *score = *score + getScore(*row, *openEnds, curPlayer)
        *row = 0
        *openEnds = 1
    } else if board[y][x] == EMPTY {
        *openEnds = 1
    } else if *row > 0 {
        *score = *score + getScore(*row, *openEnds, curPlayer)
        *row = 0
        *openEnds = 0
    } else {
        *openEnds = 0
    }
}

func checkScore(score *float64, row, openEnds *int, player int) {
    if *row > 0 {
        *score = *score + getScore(*row, *openEnds, player)
    }
    *row = 0
    *openEnds = 0
}

func currentPlayer(player1, player2 int) int {
    if player1 == player2 {
        return AI
    }
    return HUMAN
}

func getScoreFor(player1, player2, startX, endX, startY, endY int) float64 {
    score := 0.0
    row, openEnds := 0, 0

    // Horizontal
    for y := startY; y < endY; y++ {
        for x := startX; x < endX; x++ {
            checkPlayer(&score, &row, &openEnds, y, x, player1, currentPlayer(player1, player2))
        }
        checkScore(&score, &row, &openEnds, currentPlayer(player1, player2))
    }

    // Vertical
    for x := startX; x < endX; x++ {
        for y := startY; y < endY; y++ {
            checkPlayer(&score, &row, &openEnds, y, x, player1, currentPlayer(player1, player2))
        }
        checkScore(&score, &row, &openEnds, currentPlayer(player1, player2))
    }

    // Diagonal 1.1
    for tempX := startX; tempX < endX; tempX++ {
        for y, x := startY, tempX; y < endY && x < endX; y, x = y+1, x+1 {
            checkPlayer(&score, &row, &openEnds, y, x, player1, currentPlayer(player1, player2))
        }
        checkScore(&score, &row, &openEnds, currentPlayer(player1, player2))
    }

    // Diagonal 1.2
    for tempY := startY + 1; tempY < endY; tempY++ {
        for y, x := tempY, startX; y < endY && x < endX; y, x = y+1, x+1 {
            checkPlayer(&score, &row, &openEnds, y, x, player1, currentPlayer(player1, player2))
        }
        checkScore(&score, &row, &openEnds, currentPlayer(player1, player2))
    }

    // Diagonal 2.1
    for tempX := startX; tempX < endX; tempX++ {
        for y, x := startY, tempX; y < endY && x >= startX; y, x = y+1, x-1 {
            checkPlayer(&score, &row, &openEnds, y, x, player1, currentPlayer(player1, player2))
        }
        checkScore(&score, &row, &openEnds, currentPlayer(player1, player2))
    }

    // Diagonal 2.2
    for tempY := startY + 1; tempY < endY; tempY++ {
        for y, x := tempY, endX-1; y < endY && x >= startX; y, x = y+1, x-1 {
            checkPlayer(&score, &row, &openEnds, y, x, player1, currentPlayer(player1, player2))
        }
        checkScore(&score, &row, &openEnds, currentPlayer(player1, player2))
    }

    return score
}

func calculateScoreFor(pos *Position, player int) float64 {
    score1, score2 := 0., 0.
    makeMove(pos, player)
    startX, endX, startY, endY := 0, BoardWidth, 0, BoardHeight
    if pos.X > 4 {
        startX = pos.X - 4
    }
    if pos.Y > 4 {
        startY = pos.Y - 4
    }
    if pos.X < BoardWidth-5 {
        endX = pos.X + 5
    }
    if pos.Y < BoardHeight-5 {
        endY = pos.Y + 5
    }
    if pos.Capture != nil {
        score11 := getCaptureScore(player)
        score12 := getScoreFor(player, changePlayer(player), startX, endX, startY, endY)
        if score11 > score12 {
            score1 = score11
        } else {
            score1 = score12
        }
    } else {
        score1 = getScoreFor(player, changePlayer(player), startX, endX, startY, endY)
    }
    undoMove(pos, player)
    if pos.Capture != nil {
        score21 := getCaptureScore(player)
        score22 := getScoreFor(player, changePlayer(player), startX, endX, startY, endY)
        if score21 > score22 {
            score2 = score21
        } else {
            score2 = score22
        }
    } else {
        score2 = getScoreFor(player, changePlayer(player), startX, endX, startY, endY)
    }
    return score1 - score2
}

func doubleThree(pos *Position, player int) bool {
    double := false
    if pos.Capture != nil {
        return false
    }
    board[pos.Y][pos.X] = player
    if doubleThreeRule &&
        ((player == AI && freeThreeAI) || (player == HUMAN && freeThreeHuman)) &&
        checkForFreeThree(pos.Y, pos.X, player) {
        double = true
    }
    board[pos.Y][pos.X] = EMPTY
    return double
}

func getFinalScore(pos *Position, player int) (score float64) {
    if pos.Capture != nil {
        score1 := getCaptureScore(AI) - getCaptureScore(HUMAN)
        score2 := getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight) -
            getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight)

        score = getBestScore(score1, score2, player)
    } else {
        score = getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight) -
            getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight)
    }
    return
}

func checkForWin(player int) *Position {
    var win bool
    var capture *Capture = nil
    for y := 0; y < BoardHeight; y++ {
        for x := 0; x < BoardWidth; x++ {
            if board[y][x] == EMPTY && adjacentNotEmpty(y, x) {
                board[y][x] = player
                win, _ = fiveInARow(y, x, player)
                if win == true {
                    board[y][x] = EMPTY
                    return &Position{nil, y, x}
                }
                capture = finalCapture(y, x, player)
                if capture != nil {
                    return &Position{capture, y, x}
                }
                board[y][x] = EMPTY
            }
        }
    }
    return nil
}

func checkIfThereWin(player int) (win bool, pos *Position, score float64) {
    win = false

    pos = checkForWin(player)
    if pos == nil {
        pos = checkForWin(changePlayer(player))
    }

    if pos != nil {
        win = true
        makeMove(pos, player)
        score = getFinalScore(pos, player)
        undoMove(pos, player)
    }

    return
}

func generateMoves(player int) Moves {
    var moves Moves
    heap.Init(&moves)

    win, pos, score := checkIfThereWin(player)
    if win {
        return Moves{{pos, score}}
    }

    for y := 0; y < BoardHeight; y++ {
        for x := 0; x < BoardWidth; x++ {
            if board[y][x] == EMPTY && adjacentNotEmpty(y, x) {
                pos = &Position{captureMove(y, x, player), y, x}
                if doubleThree(pos, player) == false {
                    score1 := calculateScoreFor(pos, player)
                    score2 := calculateScoreFor(pos, changePlayer(player))
                    if score1 < score2 {
                        heap.Push(&moves, Move{pos, score1})
                    } else {
                        heap.Push(&moves, Move{pos, score2})
                    }
                }
            }
        }
    }

    if moves.Len() > 0 {
        last := heap.Pop(&moves).(Move)
        if last.score > 50000000 {
            return Moves{last}
        }
        heap.Push(&moves, last)
    }

    return moves
}

func changePlayer(player int) int {
    if player == AI {
        return HUMAN
    }
    return AI
}

func getBestScore(score1, score2 float64, player int) float64 {
    if (player == AI && score1 > score2) ||
        (player == HUMAN && score1 < score2) {
        return score1
    }
    return score2
}

func makeMove(pos *Position, player int) {
    board[pos.Y][pos.X] = player
    if pos.Capture != nil {
        board[pos.Capture.Pos[0].Y][pos.Capture.Pos[0].X] = EMPTY
        board[pos.Capture.Pos[1].Y][pos.Capture.Pos[1].X] = EMPTY
        captures[player] += 2
    }
}

func undoMove(pos *Position, player int) {
    if pos.Capture != nil {
        captures[player] -= 2
        board[pos.Capture.Pos[0].Y][pos.Capture.Pos[0].X] = pos.Capture.Enemy
        board[pos.Capture.Pos[1].Y][pos.Capture.Pos[1].X] = pos.Capture.Enemy
    }
    board[pos.Y][pos.X] = EMPTY
}

func minimax(player, depth int, debug *Debug) Move {
    var bestScore float64
    var pos *Position = nil
    score := 0.0

    if player == AI {
        bestScore = -1000000000
    } else {
        bestScore = 1000000000
    }

    possibleMoves := generateMoves(player)

    movesLen := possibleMoves.Len()
    for i := 0; movesLen > 0 && i < MovesCheck; movesLen, i = movesLen-1, i+1 {
        move := heap.Pop(&possibleMoves).(Move)

        makeMove(move.pos, player)
        if debugMode {
            debug.Debug = append(debug.Debug, &Debug{move.score, 0, move.pos, player, -1, []*Debug{}})
        }
        if depth == 1 {
            score = getFinalScore(move.pos, player)
        } else {
            if debugMode {
                score = minimax(changePlayer(player), depth-1, debug.Debug[i]).score
            } else {
                score = minimax(changePlayer(player), depth-1, nil).score
            }
        }
        undoMove(move.pos, player)

        bestScore = getBestScore(score, bestScore, player)
        if debugMode {
            debug.Debug[i].BestScore = bestScore
        }
        if score == bestScore {
            pos = move.pos
            if debugMode {
                debug.Index = i
            }
        }
    }

    if pos != nil && pos.Capture != nil && pos.Capture.Enemy == player {
        pos.Capture = nil
    }

    return Move{pos, bestScore}
}
