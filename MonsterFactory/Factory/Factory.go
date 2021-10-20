package Factory

import (
	"factory/MonsterFactory/Monster/Enemy"
)

type Factory interface {
	CreateOrc() Enemy.Interface
	CreateZombie() Enemy.Interface
}

//func TestFactory(factory Factory) {
//	Orc := factory.CreateOrc()
//	Orc.Say()
//
//	Zombie := factory.CreateZombie()
//	Zombie.Say()
//
//}

//func InitFactory() {
//	TestFactory(Basic{})
//	TestFactory(Hard{})
//}
