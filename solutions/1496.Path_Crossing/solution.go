package solution

import (
	"fmt"
)

func isPathCrossing(path string) bool {
	set := map[string]bool{
		"0,0": true,
	}

	p := &Point{
		X: 0,
		Y: 0,
	}

	for _, d := range path {
		handler, _ := mapper[d]
		handler(p)
		
		_, exists := set[fmt.Sprintf("%v,%v", p.X, p.Y)]
		if exists {
			return true
		}
		set[fmt.Sprintf("%v,%v", p.X, p.Y)] = true
	}

	return false
}

type Point struct {
	X int
	Y int
}

type Handler func(*Point)

func handleN(p *Point) {
	p.Y++
}

func handleS(p *Point) {
	p.Y--
}

func handleE(p *Point) {
	p.X++
}

func handleW(p *Point) {
	p.X--
}

var mapper map[rune]Handler = map[rune]Handler{
	'N': handleN,
	'S': handleS,
	'E': handleE,
	'W': handleW,
}
