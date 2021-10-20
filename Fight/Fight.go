package Fight

import (
	"factory/Board"
	"factory/MonsterFactory/Factory"
	"fmt"
)

func Fight() {
	basicFighter := Factory.Basic{}
	fighterOne := basicFighter.CreateOrc()
	fighterTwo := basicFighter.CreateZombie()

	board := Board.New(fighterOne, fighterTwo)

	fmt.Printf("----------------Приветствие----------------\n\n")
	board.BelligerentOne.Say()
	board.BelligerentTwo.Say()
	fmt.Printf("----------------БОЙ!---------------\n\n")
	for board.BelligerentOne.CanMove() && board.BelligerentTwo.CanMove() {
		board.BelligerentTwo.Beat(board.BelligerentOne)
		board.BelligerentOne.Beat(board.BelligerentTwo)
	}
	fmt.Printf("----------------бой окончен----------------\n")
}
