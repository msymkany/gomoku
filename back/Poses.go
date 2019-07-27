package main

type Poses []Pos

func (moves Poses) Len() int {
	return len(moves)
}

func (moves Poses) Less(i, j int) bool {
	return moves[i].Score > moves[j].Score
}

func (moves Poses) Swap(i, j int) {
	moves[i], moves[j] = moves[j], moves[i]
}

func (moves *Poses) Push(el interface{}) {
	*moves = append(*moves, el.(Pos))
}

func (moves *Poses) Pop() interface{} {
	item := (*moves)[moves.Len()-1]
	*moves = (*moves)[0 : moves.Len()-1]
	return item
}
