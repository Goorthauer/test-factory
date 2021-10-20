package Enemy

import (
	"factory/MonsterFactory/Magic"
	"factory/MonsterFactory/Weapon"
	"fmt"
)

type Interface interface {
	Say()
	Beat(m Interface)
	Resurrect()
	Kill()
	TakeDamage(int)
	Heal(int)
	CanMove() bool
}

type EnergyPool struct {
	Hp, MaxHp, Mana, MaxMana int
}

type Monster struct {
	Name string
	EnergyPool
	Weapon Weapon.Interface
	Magic  []Magic.Magic
	State  StateInterface
}

//Resurrect Monster
func (enemy *Monster) Resurrect() {
	enemy.State.Resurrect(enemy)
	enemy.Hp = enemy.MaxHp
}

func (enemy *Monster) setState(state StateInterface) {
	enemy.State = state
}

func (enemy *Monster) TakeDamage(amount int) {
	enemy.State.TakeDamage(enemy, amount)
}
func (enemy *Monster) Heal(amount int) {
	enemy.State.Heal(enemy, amount)
}

//Kill Monster
func (enemy *Monster) Kill() {
	enemy.State.Kill(enemy)
	enemy.Hp = 0
}

//Say Monster
func (enemy *Monster) Say() {
	enemy.beforeSay()
	enemy.State.Say(enemy)
}

// Beat Monster
func (enemy *Monster) Beat(m Interface) {
	enemy.State.Beat(enemy, m)
}

//beforeSay Monster
func (enemy *Monster) beforeSay() {
	fmt.Printf("%v говорит:	", enemy.Name)
}

func (enemy *Monster) CanMove() bool {
	return enemy.State.CanMove()
}
