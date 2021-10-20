package Zombie

import (
	"factory/MonsterFactory/Monster/Enemy"
	"factory/MonsterFactory/Weapon"
)

type SimpleZombie struct {
	Zombie
}

func CreateSimple() Enemy.Interface {
	weapon := Weapon.NewWeapon("Зубы", 1, 4)
	return New("Зомби",
		weapon,
		Enemy.EnergyPool{Hp: 25, MaxHp: 25, Mana: 8, MaxMana: 8},
		ZombieMagic())
}
