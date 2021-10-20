package Zombie

import (
	"factory/MonsterFactory/Magic"
	"factory/MonsterFactory/Monster/Enemy"
	"factory/MonsterFactory/Weapon"
)

type Zombie struct {
	Enemy.Monster
}

func ZombieMagic() []Magic.Magic {
	magicOne := Magic.NewMagic("Разложение", 8, 3, 8)
	return []Magic.Magic{*magicOne}
}

// New create Zombie
func New(name string, weapon Weapon.Interface, energy Enemy.EnergyPool, magic []Magic.Magic) Enemy.Interface {
	monster := new(Zombie)
	monster.Magic = magic
	monster.Weapon = weapon
	monster.EnergyPool = energy
	monster.Name = name
	monster.State = Enemy.AliveState{}
	return monster
}
