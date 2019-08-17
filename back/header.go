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

type Capture struct {
    Enemy int     `json:"enemy"`
    Pos   []Coord `json:"pos"`
}

type Position struct {
    Capture *Capture `json:"capture"`
    Y       int      `json:"y"`
    X       int      `json:"x"`
}

type Move struct {
    pos   *Position
    score float64
}

type Debug struct {
    TurnScore float64   `json:"turn_score"`
    BestScore float64   `json:"best_score"`
    Pos       *Position `json:"pos"`
    Player    int       `json:"player"`
    Index     int       `json:"index"`
    Debug     []*Debug  `json:"debug"`
}

// Globals

var doubleThreeRule = true
var freeThreeAI = false
var freeThreeHuman = false

var captureRule = true
var captures = map[int]int{
    AI:    0,
    HUMAN: 0,
}

var debugMode = true

var currentMove = EMPTY
var AIMode = true
var AITips = true

const maxDepth = 4
const maxMovesCheck = 150

var Depth = 3
var MovesCheck = 20

var board = [BoardHeight][BoardWidth]int{}
var win *[]Coord = nil
var winByCapture = 0
var writer http.ResponseWriter = nil
