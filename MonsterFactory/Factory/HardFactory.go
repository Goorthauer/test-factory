package Factory

import (
	"factory/MonsterFactory/Monster/Enemy"
	"factory/MonsterFactory/Monster/Orc"
	"factory/MonsterFactory/Monster/Zombie"
)

type Hard struct {
}

func (f Hard) CreateOrc() Enemy.Interface {
	return Orc.CreateWarp()
}

func (f Hard) CreateZombie() Enemy.Interface {
	return Zombie.CreateImmortal()
}
