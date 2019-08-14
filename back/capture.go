package main

func captureMove(y, x, player int) *Capture {
    if captureRule == false {
        return nil
    }

    enemy := changePlayer(player)
    poses := []int{enemy, enemy, player}

    allCoords := [][]Coord{
        {{y, x - 1}, {y, x - 2}, {y, x - 3}},
        {{y, x + 1}, {y, x + 2}, {y, x + 3}},
        {{y - 1, x}, {y - 2, x}, {y - 3, x}},
        {{y + 1, x}, {y + 2, x}, {y + 3, x}},
        {{y - 1, x - 1}, {y - 2, x - 2}, {y - 3, x - 3}},
        {{y - 1, x + 1}, {y - 2, x + 2}, {y - 3, x + 3}},
        {{y + 1, x + 1}, {y + 2, x + 2}, {y + 3, x + 3}},
        {{y + 1, x - 1}, {y + 2, x - 2}, {y + 3, x - 3}},
    }

    for i := 0; i < 8; i++ {
        if validCoordsAndPoses(allCoords[i], poses) {
            return &Capture{enemy, allCoords[i][:2]}
        }
    }
    return nil
}

func finalCapture(y, x, player int) *Capture {
    if captures[player] != 8 {
        return nil
    }

    return captureMove(y, x, player)
}

func getCaptureScore(player int) float64 {
    if captures[player] == 8 {
        return 100000000
    } else if captures[player] == 6 {
        return 10000
    } else if captures[player] == 4 {
        return 50
    } else if captures[player] == 2 {
        return 10
    } else if captures[player] == 0 {
        return 5
    }
    return 20000000
}
