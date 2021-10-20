package Enemy

import (
	"factory/MonsterFactory/Magic"
	"factory/MonsterFactory/Weapon"
	"testing"
)

func TestMonster_Heal(t *testing.T) {
	type fields struct {
		Name       string
		EnergyPool EnergyPool
		Weapon     Weapon.Weapon
		Magic      []Magic.Magic
		State      StateInterface
	}
	type args struct {
		amount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"First",
			fields{
				Name: "Первый Монстр",
				EnergyPool: EnergyPool{
					Hp:      6,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: AliveState{},
			}, args{amount: 5}},
		{"Second",
			fields{
				Name: "Второй Монстр",
				EnergyPool: EnergyPool{
					Hp:      0,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: DeadState{},
			}, args{amount: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enemy := &Monster{
				Name:       tt.fields.Name,
				EnergyPool: tt.fields.EnergyPool,
				Weapon:     tt.fields.Weapon,
				Magic:      tt.fields.Magic,
				State:      tt.fields.State,
			}
			enemy.Heal(tt.args.amount)
		})
	}
}

func TestMonster_Kill(t *testing.T) {
	type fields struct {
		Name       string
		EnergyPool EnergyPool
		Weapon     Weapon.Weapon
		Magic      []Magic.Magic
		State      StateInterface
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"First",
			fields{
				Name: "Первый Монстр",
				EnergyPool: EnergyPool{
					Hp:      6,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: AliveState{},
			}},
		{"Second",
			fields{
				Name: "Второй Монстр",
				EnergyPool: EnergyPool{
					Hp:      0,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: DeadState{},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enemy := &Monster{
				Name:       tt.fields.Name,
				EnergyPool: tt.fields.EnergyPool,
				Weapon:     tt.fields.Weapon,
				Magic:      tt.fields.Magic,
				State:      tt.fields.State,
			}
			enemy.Kill()
		})
	}
}

func TestMonster_Resurrect(t *testing.T) {
	type fields struct {
		Name       string
		EnergyPool EnergyPool
		Weapon     Weapon.Weapon
		Magic      []Magic.Magic
		State      StateInterface
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"First",
			fields{
				Name: "Первый Монстр",
				EnergyPool: EnergyPool{
					Hp:      6,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: AliveState{},
			}},
		{"Second",
			fields{
				Name: "Второй Монстр",
				EnergyPool: EnergyPool{
					Hp:      0,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: DeadState{},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enemy := &Monster{
				Name:       tt.fields.Name,
				EnergyPool: tt.fields.EnergyPool,
				Weapon:     tt.fields.Weapon,
				Magic:      tt.fields.Magic,
				State:      tt.fields.State,
			}
			enemy.Resurrect()
		})
	}
}

func TestMonster_TakeDamage(t *testing.T) {
	type fields struct {
		Name       string
		EnergyPool EnergyPool
		Weapon     Weapon.Weapon
		Magic      []Magic.Magic
		State      StateInterface
	}
	type args struct {
		amount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"First",
			fields{
				Name: "Первый Монстр",
				EnergyPool: EnergyPool{
					Hp:      6,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: AliveState{},
			}, args{amount: 4}},
		{"Second",
			fields{
				Name: "Второй Монстр",
				EnergyPool: EnergyPool{
					Hp:      3,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: AliveState{},
			}, args{amount: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enemy := &Monster{
				Name:       tt.fields.Name,
				EnergyPool: tt.fields.EnergyPool,
				Weapon:     tt.fields.Weapon,
				Magic:      tt.fields.Magic,
				State:      tt.fields.State,
			}

			enemy.TakeDamage(tt.args.amount)
		})
	}
}

func TestMonster_Cycle(t *testing.T) {
	type fields struct {
		Name       string
		EnergyPool EnergyPool
		Weapon     Weapon.Weapon
		Magic      []Magic.Magic
		State      StateInterface
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"First",
			fields{
				Name: "Первый Монстр",
				EnergyPool: EnergyPool{
					Hp:      6,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: AliveState{},
			}},
		{"Second",
			fields{
				Name: "Второй монстр",
				EnergyPool: EnergyPool{
					Hp:      0,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: DeadState{},
			}},
		{"Three",
			fields{
				Name: "Третий монстр",
				EnergyPool: EnergyPool{
					Hp:      3,
					MaxHp:   6,
					Mana:    7,
					MaxMana: 7,
				},
				Weapon: Weapon.Weapon{
					Name:   "Оружие",
					MinDmg: 1,
					MaxDmg: 2,
				},
				Magic: nil,
				State: AliveState{},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enemy := &Monster{
				Name:       tt.fields.Name,
				EnergyPool: tt.fields.EnergyPool,
				Weapon:     tt.fields.Weapon,
				Magic:      tt.fields.Magic,
				State:      tt.fields.State,
			}
			enemy.TakeDamage(3)
			enemy.Heal(2)
			enemy.TakeDamage(5)
			enemy.Kill()
			enemy.Heal(2)
			enemy.Say()
			enemy.Resurrect()
			enemy.Say()
			enemy.Resurrect()
		})
	}
}
