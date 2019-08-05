package main

import "net/http"

// Constants
const (
    EMPTY = 0
    HUMAN = 1
    AI    = 2
)

const (
    BoardWidth  = 19
    BoardHeight = 19
)

// Structures
type Coord struct {
    Y int `json:"y"`
    X int `json:"x"`
}

type Pos struct {
    Y     int
    X     int
    Score float64
}

// Globals

var doubleThreeRule = true
var freeThreeAI = false
var freeThreeHuman = false

const maxDepth = 4
const maxMovesCheck = 150

var Depth = 2
var MovesCheck = 20

var board = [BoardHeight][BoardWidth]int{}
var win *[]Coord = nil
var writer http.ResponseWriter = nil
