package main

import "net/http"

const DEPTH = 3
const MovesCheck = 10

const (
	EMPTY = 0
	HUMAN = 1
	AI    = 2
)

const (
	BoardWidth  = 19
	BoardHeight = 19
)

type Coord struct {
	Y int `json:"y"`
	X int `json:"x"`
}

type Pos struct {
	Y     int     `json:"y"`
	X     int     `json:"x"`
	Score float64 `json:"score"`
}

var board = [BoardHeight][BoardWidth]int{}
var win *[]Coord = nil
var writer http.ResponseWriter = nil
