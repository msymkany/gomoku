package main

func validCoordsAndPoses(coords []Coord, poses []int) bool {
    for index, coord := range coords {
        if coord.Y < 0 || coord.X < 0 ||
            coord.Y >= BoardHeight || coord.X >= BoardWidth ||
            board[coord.Y][coord.X] != poses[index] {
            return false
        }
    }
    return true
}

func checkForFreeThree(y, x, player int) bool {
    freeThree := false

    // Horizontal
    for tempX := x - 4; tempX <= x && freeThree == false; tempX++ {
        if validCoordsAndPoses([]Coord{
            {y, tempX}, {y, tempX + 1}, {y, tempX + 2},
            {y, tempX + 3}, {y, tempX + 4}},
            []int{EMPTY, player, player, player, EMPTY}) || validCoordsAndPoses(
            []Coord{{y, tempX}, {y, tempX + 1}, {y, tempX + 2},
                {y, tempX + 3}, {y, tempX + 4}, {y, tempX + 5}},
            []int{EMPTY, player, EMPTY, player, player, EMPTY}) || validCoordsAndPoses(
            []Coord{{y, tempX}, {y, tempX + 1}, {y, tempX + 2},
                {y, tempX + 3}, {y, tempX + 4}, {y, tempX + 5}},
            []int{EMPTY, player, player, EMPTY, player, EMPTY}) {
            freeThree = true
        }
    }

    // Vertical
    for tempY := y - 4; tempY <= y && freeThree == false; tempY++ {
        if validCoordsAndPoses([]Coord{
            {tempY, x}, {tempY + 1, x}, {tempY + 2, x},
            {tempY + 3, x}, {tempY + 4, x}},
            []int{EMPTY, player, player, player, EMPTY}) || validCoordsAndPoses(
            []Coord{{tempY, x}, {tempY + 1, x}, {tempY + 2, x},
                {tempY + 3, x}, {tempY + 4, x}, {tempY + 5, x}},
            []int{EMPTY, player, EMPTY, player, player, EMPTY}) || validCoordsAndPoses(
            []Coord{{tempY, x}, {tempY + 1, x}, {tempY + 2, x},
                {tempY + 3, x}, {tempY + 4, x}, {tempY + 5, x}},
            []int{EMPTY, player, player, EMPTY, player, EMPTY}) {
            freeThree = true
        }
    }

    // Diagonal 1
    for tempY, tempX := y-4, x-4; tempX <= x && freeThree == false; tempY, tempX = tempY+1, tempX+1 {
        if validCoordsAndPoses([]Coord{
            {tempY, tempX}, {tempY + 1, tempX + 1}, {tempY + 2, tempX + 2},
            {tempY + 3, tempX + 3}, {tempY + 4, tempX + 4}},
            []int{EMPTY, player, player, player, EMPTY}) || validCoordsAndPoses(
            []Coord{{tempY, tempX}, {tempY + 1, tempX + 1}, {tempY + 2, tempX + 2},
                {tempY + 3, tempX + 3}, {tempY + 4, tempX + 4}, {tempY + 5, tempX + 5}},
            []int{EMPTY, player, EMPTY, player, player, EMPTY}) || validCoordsAndPoses(
            []Coord{{tempY, tempX}, {tempY + 1, tempX + 1}, {tempY + 2, tempX + 2},
                {tempY + 3, tempX + 3}, {tempY + 4, tempX + 4}, {tempY + 5, tempX + 5}},
            []int{EMPTY, player, player, EMPTY, player, EMPTY}) {
            freeThree = true
        }
    }

    // Diagonal 2
    for tempY, tempX := y+4, x-4; tempX <= x && freeThree == false; tempY, tempX = tempY-1, tempX+1 {
        if validCoordsAndPoses([]Coord{
            {tempY, tempX}, {tempY - 1, tempX + 1}, {tempY - 2, tempX + 2},
            {tempY - 3, tempX + 3}, {tempY - 4, tempX + 4}},
            []int{EMPTY, player, player, player, EMPTY}) || validCoordsAndPoses(
            []Coord{{tempY, tempX}, {tempY - 1, tempX + 1}, {tempY - 2, tempX + 2},
                {tempY - 3, tempX + 3}, {tempY - 4, tempX + 4}, {tempY - 5, tempX + 5}},
            []int{EMPTY, player, EMPTY, player, player, EMPTY}) || validCoordsAndPoses(
            []Coord{{tempY, tempX}, {tempY - 1, tempX + 1}, {tempY - 2, tempX + 2},
                {tempY - 3, tempX + 3}, {tempY - 4, tempX + 4}, {tempY - 5, tempX + 5}},
            []int{EMPTY, player, player, EMPTY, player, EMPTY}) {
            freeThree = true
        }
    }

    return freeThree
}

func resetFreeThrees() {
    freeThreeAI = false
    freeThreeHuman = false

    for y := 0; y < BoardHeight; y++ {
        for x := 0; x < BoardWidth; x++ {
            if freeThreeAI == true && freeThreeHuman == true {
                return
            }

            if freeThreeAI == false && board[y][x] == AI &&
                checkForFreeThree(y, x, AI) {
                freeThreeAI = true
            }

            if freeThreeHuman == false && board[y][x] == HUMAN &&
                checkForFreeThree(y, x, HUMAN) {
                freeThreeHuman = true
            }
        }
    }

}
