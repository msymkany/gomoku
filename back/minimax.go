package main

import "container/heap"

func fiveInARow(y, x, player int) (bool, *[]Coord) {
    inARow := 0
    line := [9]Coord{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
    var tempX, tempY int

    // Horizontal
    for tempX = x - 4; tempX <= x+4 && tempX < BoardWidth; tempX++ {
        if inARow >= 5 && tempX >= 0 && tempX < BoardWidth && board[y][tempX] != player {
            break
        }
        if tempX >= 0 && tempX < BoardWidth && board[y][tempX] == player {
            line[inARow].Y, line[inARow].X = y, tempX
            inARow++
        } else {
            inARow = 0
        }
    }
    if inARow < 5 {
        inARow = 0
    } else {
        winLine := line[:inARow]
        return true, &winLine
    }

    // Vertical
    for tempY = y - 4; tempY <= y+4 && tempY < BoardHeight; tempY++ {
        if inARow >= 5 && tempY >= 0 && tempY < BoardHeight && board[tempY][x] != player {
            break
        }
        if tempY >= 0 && tempY < BoardHeight && board[tempY][x] == player {
            line[inARow].Y, line[inARow].X = tempY, x
            inARow++
        } else {
            inARow = 0
        }
    }
    if inARow < 5 {
        inARow = 0
    } else {
        winLine := line[:inARow]
        return true, &winLine
    }

    // Diagonal 1
    for tempY, tempX = y-4, x-4; tempX <= x+4 && tempY < BoardHeight && tempX < BoardWidth; tempY, tempX = tempY+1, tempX+1 {
        if inARow >= 5 && tempY >= 0 && tempY < BoardHeight && tempX >= 0 && tempX < BoardWidth &&
            board[tempY][tempX] != player {
            break
        }
        if tempY >= 0 && tempY < BoardHeight && tempX >= 0 && tempX < BoardWidth && board[tempY][tempX] == player {
            line[inARow].Y, line[inARow].X = tempY, tempX
            inARow++
        } else {
            inARow = 0
        }
    }
    if inARow < 5 {
        inARow = 0
    } else {
        winLine := line[:inARow]
        return true, &winLine
    }

    // Diagonal 2
    for tempY, tempX = y+4, x-4; tempX <= x+4 && tempY >= 0 && tempX < BoardWidth; tempY, tempX = tempY-1, tempX+1 {
        if inARow >= 5 && tempY >= 0 && tempY < BoardHeight && tempX >= 0 && tempX < BoardWidth &&
            board[tempY][tempX] != player {
            break
        }
        if tempY >= 0 && tempY < BoardHeight && tempX >= 0 && tempX < BoardWidth && board[tempY][tempX] == player {
            line[inARow].Y, line[inARow].X = tempY, tempX
            inARow++
        } else {
            inARow = 0
        }
    }
    if inARow < 5 {
        inARow = 0
    } else {
        winLine := line[:inARow]
        return true, &winLine
    }
    winLine := line[:inARow]

    return false, &winLine
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

func currentPlayer(player1, player2 int) int {
    if player1 == player2 {
        return AI
    }
    return HUMAN
}

func getScoreFor(player1, player2, startX, endX, startY, endY int) float64 {
    var x, y, tempX, tempY int
    score := 0.0
    row, openEnds := 0, 0

    // Horizontal
    for y = startY; y < endY; y++ {
        for x = startX; x < endX; x++ {
            if board[y][x] == player1 {
                row = row + 1
            } else if board[y][x] == EMPTY && row > 0 {
                openEnds = openEnds + 1
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 1
            } else if board[y][x] == EMPTY {
                openEnds = 1
            } else if row > 0 {
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 0
            } else {
                openEnds = 0
            }
        }
        if row > 0 {
            score = score + getScore(row, openEnds, currentPlayer(player1, player2))
        }
        row = 0
        openEnds = 0
    }

    // Vertical
    for x = startX; x < endX; x++ {
        for y = startY; y < endY; y++ {
            if board[y][x] == player1 {
                row = row + 1
            } else if board[y][x] == EMPTY && row > 0 {
                openEnds = openEnds + 1
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 1
            } else if board[y][x] == EMPTY {
                openEnds = 1
            } else if row > 0 {
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 0
            } else {
                openEnds = 0
            }
        }
        if row > 0 {
            score = score + getScore(row, openEnds, currentPlayer(player1, player2))
        }
        row = 0
        openEnds = 0
    }

    // Diagonal 1.1
    for tempX = startX; tempX < endX; tempX++ {
        for y, x = startY, tempX; y < endY && x < endX; y, x = y+1, x+1 {
            if board[y][x] == player1 {
                row = row + 1
            } else if board[y][x] == EMPTY && row > 0 {
                openEnds = openEnds + 1
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 1
            } else if board[y][x] == EMPTY {
                openEnds = 1
            } else if row > 0 {
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 0
            } else {
                openEnds = 0
            }
        }
        if row > 0 {
            score = score + getScore(row, openEnds, currentPlayer(player1, player2))
        }
        row = 0
        openEnds = 0
    }

    // Diagonal 1.2
    for tempY = startY + 1; tempY < endY; tempY++ {
        for y, x = tempY, startX; y < endY && x < endX; y, x = y+1, x+1 {
            if board[y][x] == player1 {
                row = row + 1
            } else if board[y][x] == EMPTY && row > 0 {
                openEnds = openEnds + 1
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 1
            } else if board[y][x] == EMPTY {
                openEnds = 1
            } else if row > 0 {
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 0
            } else {
                openEnds = 0
            }
        }
        if row > 0 {
            score = score + getScore(row, openEnds, currentPlayer(player1, player2))
        }
        row = 0
        openEnds = 0
    }

    // Diagonal 2.1
    for tempX = startX; tempX < endX; tempX++ {
        for y, x = startY, tempX; y < endY && x >= startX; y, x = y+1, x-1 {
            if board[y][x] == player1 {
                row = row + 1
            } else if board[y][x] == EMPTY && row > 0 {
                openEnds = openEnds + 1
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 1
            } else if board[y][x] == EMPTY {
                openEnds = 1
            } else if row > 0 {
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 0
            } else {
                openEnds = 0
            }
        }
        if row > 0 {
            score = score + getScore(row, openEnds, currentPlayer(player1, player2))
        }
        row = 0
        openEnds = 0
    }

    // Diagonal 2.2
    for tempY = startY + 1; tempY < endY; tempY++ {
        for y, x = tempY, endX-1; y < endY && x >= startX; y, x = y+1, x-1 {
            if board[y][x] == player1 {
                row = row + 1
            } else if board[y][x] == EMPTY && row > 0 {
                openEnds = openEnds + 1
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 1
            } else if board[y][x] == EMPTY {
                openEnds = 1
            } else if row > 0 {
                score = score + getScore(row, openEnds, currentPlayer(player1, player2))
                row = 0
                openEnds = 0
            } else {
                openEnds = 0
            }
        }
        if row > 0 {
            score = score + getScore(row, openEnds, currentPlayer(player1, player2))
        }
        row = 0
        openEnds = 0
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
        if pos.Capture != nil {
            score1 := getCaptureScore(AI) - getCaptureScore(HUMAN)
            score2 := getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight) -
                getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight)
            if (player == AI && score1 > score2) ||
                (player == HUMAN && score1 < score2) {
                score = score1
            } else {
                score = score2
            }
        } else {
            score = getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight) -
                getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight)
        }
        undoMove(pos, player)
    }

    return
}

func generateMoves(player int) Moves {
    var moves Moves
    var score1, score2 float64
    //var score float64
    heap.Init(&moves)

    if win, pos, score := checkIfThereWin(player); win {
        return Moves{{pos, score}}
    }

    for y := 0; y < BoardHeight; y++ {
        for x := 0; x < BoardWidth; x++ {
            if board[y][x] == EMPTY && adjacentNotEmpty(y, x) {
                pos := &Position{captureMove(y, x, player), y, x}
                if doubleThree(pos, player) == false {
                    score1 = calculateScoreFor(pos, player)
                    score2 = calculateScoreFor(&Position{captureMove(y, x, changePlayer(player)), y, x}, changePlayer(player))
                    if score1 < score2 {
                        heap.Push(&moves, Move{pos, score1})
                    } else {
                        heap.Push(&moves, Move{pos, score2})
                    }
                    //            if pos.Capture != nil {
                    //                score = getBestScore(getCaptureScore(AI)-getCaptureScore(HUMAN),
                    //                    getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight)-
                    //                        getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight), player)
                    //            } else {
                    //                score = getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight) -
                    //                    getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight)
                    //            }
                    //heap.Push(&moves, Move{pos, score})

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

func minimax(player, depth int, debug *Debug, alpha, beta float64) Move {
    var bestScore float64
    var pos *Position = nil
    var move Move
    score := 0.0

    if player == AI {
        bestScore = -1000000000
    } else {
        bestScore = 1000000000
    }

    possibleMoves := generateMoves(player)

    movesLen := possibleMoves.Len()
    for i := 0; movesLen > 0 && i < MovesCheck; movesLen, i = movesLen-1, i+1 {
        move = heap.Pop(&possibleMoves).(Move)

        makeMove(move.pos, player)
        if debugMode {
            debug.Debug = append(debug.Debug, &Debug{move.score, 0, move.pos, player, -1, []*Debug{}})
        }
        if depth == 1 {
            //            score = move.score
            if move.pos.Capture != nil {
                score1 := getCaptureScore(AI) - getCaptureScore(HUMAN)
                score2 := getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight) -
                    getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight)
                if (player == AI && score1 > score2) ||
                    (player == HUMAN && score1 < score2) {
                    score = score1
                } else {
                    score = score2
                }
            } else {
                score = getScoreFor(AI, changePlayer(player), 0, BoardWidth, 0, BoardHeight) -
                    getScoreFor(HUMAN, changePlayer(player), 0, BoardWidth, 0, BoardHeight)
            }
        } else {
            if debugMode {
                score = minimax(changePlayer(player), depth-1, debug.Debug[i], alpha, beta).score
            } else {
                score = minimax(changePlayer(player), depth-1, nil, alpha, beta).score
            }
        }
        undoMove(move.pos, player)

        if (player == AI && score > bestScore) ||
            (player == HUMAN && score < bestScore) {
            bestScore = score
            pos = move.pos
            if debugMode {
                debug.Index = i
            }
        }
        if debugMode {
            debug.Debug[i].BestScore = bestScore
        }

        if player == AI && bestScore > alpha {
            alpha = bestScore
        } else if player == HUMAN && bestScore < beta {
            beta = bestScore
        }

        if beta <= alpha {
            break
        }
    }

    if pos != nil && pos.Capture != nil && pos.Capture.Enemy == player {
        pos.Capture = nil
    }

    return Move{pos, bestScore}
}
