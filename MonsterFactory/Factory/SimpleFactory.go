package Factory

import (
	"factory/MonsterFactory/Monster/Enemy"
	"factory/MonsterFactory/Monster/Orc"
	"factory/MonsterFactory/Monster/Zombie"
)

type Basic struct {
}

func (f Basic) CreateOrc() Enemy.Interface {
	return Orc.CreateSimple()
}

func (f Basic) CreateZombie() Enemy.Interface {
	return Zombie.CreateSimple()
}
