package Zombie

import (
	"factory/MonsterFactory/Monster/Enemy"
	"factory/MonsterFactory/Weapon"
)

type ImmortalZombie struct {
	Zombie
}

func CreateImmortal() Enemy.Interface {
	weapon := Weapon.NewWeapon("Меч", 2, 5)
	return New("Бессмертный зомби",
		weapon,
		Enemy.EnergyPool{Hp: 7, MaxHp: 7, Mana: 15, MaxMana: 15},
		ZombieMagic())
}
