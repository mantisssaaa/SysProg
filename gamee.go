package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character interface {
	Hit() string
	Block() string
	GetHP() int
	GetName() string
	TakeDamage(damage int)
}

type Player struct {
	Name      string
	HP        int
	Strength  int
	hitPart   string
	blockPart string
}

type Enemy struct {
	Name      string
	HP        int
	Strength  int
	hitPart   string
	blockPart string
}

func (p *Player) GetHP() int {
	return p.HP
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) Hit() string {
	return p.hitPart
}

func (p *Player) Block() string {
	return p.blockPart
}

func (p *Player) TakeDamage(damage int) {
	p.HP -= damage
	if p.HP < 0 {
		p.HP = 0
	}
}

func (e *Enemy) GetHP() int {
	return e.HP
}

func (e *Enemy) GetName() string {
	return e.Name
}

func (e *Enemy) Hit() string {
	return e.hitPart
}

func (e *Enemy) Block() string {
	return e.blockPart
}

func (e *Enemy) TakeDamage(damage int) {
	e.HP -= damage
	if e.HP < 0 {
		e.HP = 0
	}
}

func playerChoice() (string, string) {
	var hit, block string

	fmt.Println("\nВыберите часть тела для атаки (голова/грудь/ноги):")
	fmt.Scan(&hit)

	fmt.Println("Выберите часть тела для защиты (голова/грудь/ноги):")
	fmt.Scan(&block)

	return hit, block
}

func computerChoice() (string, string) {
	parts := []string{"голова", "грудь", "ноги"}

	hitIndex := rand.Intn(len(parts))
	blockIndex := rand.Intn(len(parts))

	return parts[hitIndex], parts[blockIndex]
}

func fight(player Character, enemy Character) {
	round := 1

	var cheat string
	for player.GetHP() > 0 && enemy.GetHP() > 0 {

		fmt.Printf("\n=== Раунд %d ===\n", round)

		playerObj := player.(*Player)
		enemyObj := enemy.(*Enemy)

		if cheat == " " {
			enemy.TakeDamage(1000)
			fmt.Printf("Победитель: %s\n", player.GetName())
			fmt.Printf("Проигравший: %s\n", enemy.GetName())
			return
		}

		playerObj.hitPart, playerObj.blockPart = playerChoice()

		enemyObj.hitPart, enemyObj.blockPart = computerChoice()

		fmt.Printf("Игрок атакует: %s, защищает: %s\n", playerObj.hitPart, playerObj.blockPart)
		fmt.Printf("Противник атакует: %s, защищает: %s\n", enemyObj.hitPart, enemyObj.blockPart)

		if player.Hit() != enemy.Block() {
			enemy.TakeDamage(playerObj.Strength)
			fmt.Printf("Игрок попал! У противника осталось %d HP\n", enemy.GetHP())
		} else {
			fmt.Println("Противник заблокировал удар игрока!")
		}

		if player.GetHP() > 0 {
			if enemy.Hit() != player.Block() {
				player.TakeDamage(enemyObj.Strength)
				fmt.Printf("Противник попал! У игрока осталось %d HP\n", player.GetHP())
			} else {
				fmt.Println("Игрок заблокировал удар противника!")
			}
		}

		round++

	}

	fmt.Println("\n=== Бой окончен! ===")
	if player.GetHP() > 0 {
		fmt.Printf("Победитель: %s\n", player.GetName())
		fmt.Printf("Проигравший: %s\n", enemy.GetName())
	} else {
		fmt.Printf("Победитель: %s\n", enemy.GetName())
		fmt.Printf("Проигравший: %s\n", player.GetName())
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Начало игры ===")

	player := &Player{
		Name:     "Герой",
		HP:       100,
		Strength: 20,
	}

	enemy := &Enemy{
		Name:     "Злодей",
		HP:       80,
		Strength: 15,
	}

	//// Запуск боя \\\\
	fight(player, enemy)
}
