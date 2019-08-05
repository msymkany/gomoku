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

func checkForWin(player int) (int, int, bool) {
    var win bool
    for y := 0; y < BoardHeight; y++ {
        for x := 0; x < BoardWidth; x++ {
            if board[y][x] == EMPTY && adjacentNotEmpty(y, x) {
                board[y][x] = player
                win, _ = fiveInARow(y, x, player)
                if win == true {
                    board[y][x] = EMPTY
                    return y, x, true
                }
                board[y][x] = EMPTY
            }
        }
    }
    return 0, 0, false
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

func calculateScoreFor(player, y, x int) float64 {
    board[y][x] = player
    startX, endX, startY, endY := 0, BoardWidth, 0, BoardHeight
    if x > 4 {
        startX = x - 4
    }
    if y > 4 {
        startY = y - 4
    }
    if x < BoardWidth-5 {
        endX = x + 5
    }
    if y < BoardHeight-5 {
        endY = y + 5
    }
    score := getScoreFor(player, changePlayer(player), startX, endX, startY, endY)
    board[y][x] = EMPTY
    return score - getScoreFor(player, changePlayer(player), startX, endX, startY, endY)
}

func doubleThree(y, x, player int) bool {
    double := false
    board[y][x] = player
    if doubleThreeRule &&
        ((player == AI && freeThreeAI) || (player == HUMAN && freeThreeHuman)) &&
        checkForFreeThree(y, x, player) {
        double = true
    }
    board[y][x] = EMPTY
    return double
}

func generateMoves(player int) Poses {
    var moves Poses
    heap.Init(&moves)

    y, x, isWin := checkForWin(player)
    if isWin == false {
        y, x, isWin = checkForWin(changePlayer(player))
    }
    if isWin == true {
        board[y][x] = player
        score := getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight) -
            getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight)
        board[y][x] = EMPTY
        return Poses{{y, x, score}}
    }

    for y := 0; y < BoardHeight; y++ {
        for x := 0; x < BoardWidth; x++ {
            if board[y][x] == EMPTY && adjacentNotEmpty(y, x) &&
                doubleThree(y, x, player) == false {
                score1 := calculateScoreFor(player, y, x)
                score2 := calculateScoreFor(changePlayer(player), y, x)
                if score1 < score2 {
                    heap.Push(&moves, Pos{y, x, score1})
                } else {
                    heap.Push(&moves, Pos{y, x, score2})
                }
            }
        }
    }

    if moves.Len() > 0 {
        last := heap.Pop(&moves).(Pos)
        if last.Score > 50000000 {
            return Poses{last}
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

func minimax(player, depth int) Pos {
    var bestScore float64
    y, x, score := -1, -1, 0.0

    if player == AI {
        bestScore = -1000000000
    } else {
        bestScore = 1000000000
    }

    possibleMoves := generateMoves(player)

    movesLen := possibleMoves.Len()
    for i := 0; movesLen > 0 && i < MovesCheck; movesLen, i = movesLen-1, i+1 {
        move := heap.Pop(&possibleMoves).(Pos)

        board[move.Y][move.X] = player
        if depth == 1 {
            score = getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight) -
                getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight)
        } else {
            score = minimax(changePlayer(player), depth-1).Score
        }
        board[move.Y][move.X] = EMPTY

        if (score > bestScore && player == AI) ||
            (score < bestScore && player == HUMAN) {
            bestScore, y, x = score, move.Y, move.X
        }
    }

    return Pos{y, x, bestScore}
}
