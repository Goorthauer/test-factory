package Enemy

import (
	"fmt"
	"math/rand"
	"strconv"
)

// StateInterface
//@todo Нужно подумать как тут всё корректно разнести
type StateInterface interface {
	CanMove() bool
	Resurrect(c *Monster)
	Kill(c *Monster)
	Say(c *Monster)
	Beat(c *Monster, m Interface)
	TakeDamage(c *Monster, amount int)
	Heal(c *Monster, amount int)
}

//DeadState Мертв
type DeadState struct{}

func (ds DeadState) CanMove() bool {
	return false
}
func (ds DeadState) Resurrect(m *Monster) {
	fmt.Printf("%v воскрешен.\n", m.Name)
	m.setState(AliveState{})
}
func (ds DeadState) Kill(m *Monster) {
	fmt.Printf("Мертвого %v убить нельзя\n", m.Name)
}
func (ds DeadState) Say(m *Monster) {
	fmt.Printf("Мертвый %v не может говорить\n", m.Name)
}
func (ds DeadState) Beat(c *Monster, _ Interface) {
	fmt.Printf("Мертвый %v не может наносить удары\n", c.Name)
}

func (ds DeadState) TakeDamage(m *Monster, _ int) {
	fmt.Printf("Монстр %v Мёртв, ему нельзя нанести урон\n", m.Name)
}
func (ds DeadState) Heal(m *Monster, _ int) {
	fmt.Printf("%v мёртв, его нельзя лечить\n", m.Name)
}

//AliveState Жив
type AliveState struct{}

func (a AliveState) Resurrect(m *Monster) {
	fmt.Printf("Живого %v невозможно воскресить\n", m.Name)
}
func (a AliveState) Kill(m *Monster) {
	fmt.Printf("%v убит\n", m.Name)
	m.setState(DeadState{})
}
func (a AliveState) Say(m *Monster) {
	fmt.Printf("А А А!1 Я имею %v Жизней и %v маны\n", m.Hp, m.Mana)

}
func (a AliveState) Beat(c *Monster, m Interface) {
	damage := c.Weapon.GetDamage()
	weapon := c.Weapon.GetName()
	//@todo Перебрать магию, сделать так что бы она кастовалась и убирала ману через интерфейсы
	if len(c.Magic) > 0 && c.Mana > 0 {
		magic := c.Magic[rand.Intn(len(c.Magic))]
		magicDamage := rand.Intn(magic.MaxDmg-magic.MinDmg) + magic.MinDmg
		if magicDamage > 0 && c.Mana-magic.Mana >= 0 {
			damage = magicDamage
			c.Mana = c.Mana - magic.Mana
			weapon = magic.Name + ". Израсходованно маны: " + strconv.Itoa(magic.Mana)
		}
	}
	fmt.Printf("Монстр %v ударил на %v единиц урона(%v).\n", c.Name, damage, weapon)
	m.TakeDamage(damage)

}

func (a AliveState) TakeDamage(m *Monster, amount int) {
	m.Hp = m.Hp - amount
	if m.Hp <= 0 {
		m.Hp = 0
		m.setState(DeadState{})
		fmt.Printf("Монстр %v убит. ", m.Name)
	}
	fmt.Printf("У монстра %v осталось %v единиц здоровья\n", m.Name, m.Hp)
}
func (a AliveState) Heal(m *Monster, amount int) {
	m.Hp = m.Hp + amount
	if m.Hp > m.MaxHp {
		m.Hp = m.MaxHp
	}
	fmt.Printf("%v вылечен на %v единиц здоровья\n", m.Name, m.Hp)
}

func (a AliveState) CanMove() bool {
	return true
}
