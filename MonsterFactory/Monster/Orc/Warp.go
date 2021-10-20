package Orc

import (
	"factory/MonsterFactory/Magic"
	"factory/MonsterFactory/Monster/Enemy"
	"factory/MonsterFactory/Weapon"
)

type WarpOrc struct {
	Orc
}

func WarpOrcMagic() []Magic.Magic {
	magicOne := Magic.NewMagic("ВААГХ!11", 5, 1, 2)
	magicTwo := Magic.NewMagic("РЁВ Вожака!", 10, 3, 5)
	return []Magic.Magic{*magicOne, *magicTwo}
}
func CreateWarp() Enemy.Interface {
	weapon := Weapon.NewWeapon("Убийца драконов", 3, 5)
	magics := WarpOrcMagic()
	return New("Вожак орков",
		weapon,
		Enemy.EnergyPool{Hp: 10, MaxHp: 10, Mana: 15, MaxMana: 15},
		magics)
}
