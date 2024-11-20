package main

import (
	"fmt"
)

type player struct{ hp, mana, armor int }
type boss struct{ hp, damage int }

var spells = []spell{
	{53, func(p *player, b *boss) { b.hp -= 4 }, effect{}},
	{73, func(p *player, b *boss) { p.hp += 2; b.hp -= 2 }, effect{}},
	{113, func(p *player, b *boss) { p.armor += 7 }, effect{6, func(p *player, b *boss, timer int) {
		if timer == 0 {
			p.armor -= 7
		}
	}}},
	{173, func(*player, *boss) {}, effect{6, func(p *player, b *boss, timer int) { b.hp -= 3 }}},
	{229, func(*player, *boss) {}, effect{5, func(p *player, b *boss, timer int) { p.mana += 101 }}},
}

var spellNames = []string{"Missile", "Drain", "Shield", "Poison", "Recharge"}

type spell struct {
	mana    int
	instant instant
	effect  effect
}

type instant func(*player, *boss)

type effect struct {
	timer int
	f     func(*player, *boss, int)
}

// return value reports if the effect should still be active.
func (e *effect) apply(p *player, b *boss) bool {
	if e.timer <= 0 {
		return false
	}
	e.timer--
	e.f(p, b, e.timer)
	return e.timer > 0
}

func main() {
	p := player{50, 500, 0}
	b := boss{71, 10}
	//p = player{10, 250, 0}
	//b = boss{13, 8}
	minimize(p, b, make(map[int]*effect), 0, nil)
	fmt.Println("Done\n")
	fmt.Println("part1:", mincost)
	fmt.Println(best)
	effects := make(map[int]*effect)
	for _, i := range best {
		sp := spells[i]

		fmt.Println("Player")
		fmt.Println(p, b)
		for i, eff := range effects {
			fmt.Println("Effect", spellNames[i])
			if !eff.apply(&p, &b) {
				delete(effects, i)
			}
		}
		fmt.Println(p, b)
		fmt.Println("Cast", spellNames[i])
		p.mana -= sp.mana
		sp.instant(&p, &b)
		effects[i] = &sp.effect
		fmt.Println(p, b)
		fmt.Println()

		fmt.Println("Boss")
		fmt.Println(p, b)
		for i, eff := range effects {
			fmt.Println("Effect", spellNames[i])
			if !eff.apply(&p, &b) {
				delete(effects, i)
			}
		}
		fmt.Println(p, b)
		p.hp -= max(1, b.damage-p.armor)
		fmt.Println(p, b)
		fmt.Println()
	}
	fmt.Println(p, b)
}

var mincost = 9999999999

var best []int

func minimize(p player, b boss, effects map[int]*effect, cost int, casted []int) {
	if cost >= mincost {
		return
	}
	// player's turn
	for i, eff := range effects {
		if !eff.apply(&p, &b) {
			delete(effects, i)
		}
	}
	for i, sp := range spells {
		if p.mana < sp.mana {
			continue
		}
		if _, ok := effects[i]; ok {
			// can't have multiple of the same effect active
			continue
		}
		castedCopy := make([]int, len(casted)+1)
		copy(castedCopy, casted)
		castedCopy[len(casted)] = i
		p, b, effects := p, b, clone(effects)
		p.mana -= sp.mana
		sp.instant(&p, &b)
		effects[i] = &sp.effect
		// boss's turn
		for i, eff := range effects {
			if !eff.apply(&p, &b) {
				delete(effects, i)
			}
		}
		if b.hp <= 0 {
			if cost+sp.mana < mincost {
				mincost = cost + sp.mana
				best = castedCopy
			}
			continue
		}
		p.hp -= max(1, b.damage-p.armor)
		if p.hp <= 0 {
			continue
		}
		minimize(p, b, effects, cost+sp.mana, castedCopy)
	}
}

func clone(effects map[int]*effect) map[int]*effect {
	out := make(map[int]*effect)
	for i, eff := range effects {
		effCopy := *eff
		out[i] = &effCopy
	}
	return out
}
