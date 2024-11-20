package main

import (
	"fmt"
	"math"
)

type player struct{ hp, mana, armor int }
type boss struct{ hp, damage int }

type spell struct {
	mana    int
	instant func(*player, *boss)
	effect  effect
}

type effect struct {
	timer int
	f     func(*player, *boss, int)
}

type effectMap map[int]*effect

func (effects effectMap) apply(p *player, b *boss) {
	for i, eff := range effects {
		if eff.timer <= 0 {
			delete(effects, i)
			continue
		}
		eff.timer--
		eff.f(p, b, eff.timer)
		if eff.timer <= 0 {
			delete(effects, i)
		}
	}
}

func (effects effectMap) clone() effectMap {
	out := make(effectMap)
	for i, eff := range effects {
		effCopy := *eff
		out[i] = &effCopy
	}
	return out
}

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

func main() {
	p := player{50, 500, 0}
	b := boss{71, 10}
	minimize(p, b, make(effectMap), 0)
	fmt.Println("part1:", mincost)
}

var mincost = math.MaxInt

func minimize(p player, b boss, effects effectMap, cost int) {
	if cost >= mincost {
		return
	}
	// Player's turn
	effects.apply(&p, &b)
	for i, sp := range spells {
		if p.mana < sp.mana {
			continue
		}
		if _, ok := effects[i]; ok {
			// Can't have multiple of the same effect active.
			continue
		}
		// Copy state to not affect previous/next iteration.
		p, b, effects := p, b, effects.clone()
		p.mana -= sp.mana
		sp.instant(&p, &b)
		effects[i] = &sp.effect
		// Boss's turn
		effects.apply(&p, &b)
		if b.hp <= 0 {
			mincost = min(mincost, cost+sp.mana)
			continue
		}
		p.hp -= max(1, b.damage-p.armor)
		if p.hp <= 0 {
			continue
		}
		minimize(p, b, effects, cost+sp.mana)
	}
}
