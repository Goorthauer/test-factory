package Orc

import (
	"factory/MonsterFactory/Magic"
	"factory/MonsterFactory/Monster/Enemy"
	"factory/MonsterFactory/Weapon"
)

type Orc struct {
	Enemy.Monster
}

// New create Orc
func New(name string, weapon Weapon.Interface, energy Enemy.EnergyPool, magic []Magic.Magic) Enemy.Interface {
	monster := new(Orc)
	monster.Magic = magic
	monster.Weapon = weapon
	monster.EnergyPool = energy
	monster.Name = name
	monster.State = Enemy.AliveState{}
	return monster
}
