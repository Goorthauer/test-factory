package Weapon

import "math/rand"

type Weapon struct {
	Name           string
	MinDmg, MaxDmg int
}

type Interface interface {
	GetDamage() int
	GetName() string
}

func NewWeapon(name string, minDmg int, maxDmg int) Interface {
	item := new(Weapon)
	item.Name = name
	item.MaxDmg = maxDmg
	item.MinDmg = minDmg
	return item
}

func (w *Weapon) GetDamage() int {
	return rand.Intn(w.MaxDmg-w.MinDmg) + w.MinDmg
}

func (w *Weapon) GetName() string {
	return w.Name
}
