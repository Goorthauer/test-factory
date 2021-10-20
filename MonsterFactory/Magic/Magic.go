package Magic

type Magic struct {
	Name                 string
	Mana, MinDmg, MaxDmg int
}

func NewMagic(name string, mana int, minDmg int, maxDmg int) *Magic {
	spell := new(Magic)
	spell.Name = name
	spell.Mana = mana
	spell.MinDmg = minDmg
	spell.MaxDmg = maxDmg
	return spell
}
