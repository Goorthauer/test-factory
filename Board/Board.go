package Board

import "factory/MonsterFactory/Monster/Enemy"

type Board struct {
	BelligerentOne Enemy.Interface
	BelligerentTwo Enemy.Interface
}

func New(belligerentOne Enemy.Interface, belligerentTwo Enemy.Interface) *Board {
	board := new(Board)
	board.BelligerentTwo = belligerentTwo
	board.BelligerentOne = belligerentOne
	return board
}
