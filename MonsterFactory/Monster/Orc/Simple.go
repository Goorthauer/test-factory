package Orc

import (
	"factory/MonsterFactory/Magic"
	"factory/MonsterFactory/Monster/Enemy"
	"factory/MonsterFactory/Weapon"
)

type SimpleOrc struct {
	Orc
}

func CreateSimple() Enemy.Interface {
	weapon := Weapon.NewWeapon("Топор", 2, 4)
	return New("Орк",
		weapon,
		Enemy.EnergyPool{Hp: 25, MaxHp: 25, Mana: 10, MaxMana: 10}, []Magic.Magic{})
}
