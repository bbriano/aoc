package main

import (
	"fmt"
	"math"
)

type stat struct{ hp, damage, armor int }

var none = stat{0, 0, 0}

var weapons = []stat{
	stat{8, 4, 0},
	stat{10, 5, 0},
	stat{25, 6, 0},
	stat{40, 7, 0},
	stat{74, 8, 0},
}

var armors = []stat{
	none,
	stat{13, 0, 1},
	stat{31, 0, 2},
	stat{53, 0, 3},
	stat{75, 0, 4},
	stat{102, 0, 5},
}

var rings = []stat{
	none,
	stat{25, 1, 0},
	stat{50, 2, 0},
	stat{100, 3, 0},
	stat{20, 0, 1},
	stat{40, 0, 2},
	stat{80, 0, 3},
}

func main() {
	mincost, maxcost := math.MaxInt, 0
	for _, w := range weapons {
		for _, a := range armors {
			for _, r1 := range rings {
				for _, r2 := range rings {
					if r1 == r2 && r2 != none {
						continue
					}
					boss := stat{100, 8, 2}
					player := stat{100, 0, 0}
					cost := 0
					for _, s := range []stat{w, a, r1, r2} {
						player.damage += s.damage
						player.armor += s.armor
						cost += s.hp
					}
					for {
						boss.hp -= max(1, player.damage-boss.armor)
						if boss.hp <= 0 {
							break
						}
						player.hp -= max(1, boss.damage-player.armor)
					}
					if player.hp > 0 && cost < mincost {
						mincost = min(mincost, cost)
					}
					if player.hp <= 0 && cost > maxcost {
						maxcost = max(maxcost, cost)
					}
				}
			}
		}
	}
	fmt.Println("part1:", mincost)
	fmt.Println("part2:", maxcost)
}
