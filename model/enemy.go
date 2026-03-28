package model

import (
	"math/rand"

	"github.com/probeldev/gameshooter/config"
)

type Enemy struct {
	X int
	Y int
}

func NewEnemy() Enemy {
	e := Enemy{}
	e.setSpawnPosition()

	return e
}

func (e *Enemy) setSpawnPosition() {
	positions := []string{
		"left",
		"right",
		"top",
		"bottom",
	}

	position := positions[rand.Intn(len(positions)-1)]

	switch position {
	case "left":
		e.X = 0
		e.Y = rand.Intn(config.CountPointY - 1)
	case "right":
		e.X = config.CountPointX - 1
		e.Y = rand.Intn(config.CountPointY - 1)
	case "top":
		e.X = rand.Intn(config.CountPointX - 1)
		e.Y = 0
	case "bottom":
		e.X = rand.Intn(config.CountPointX - 1)
		e.Y = config.CountPointX - 1
	}
}

func (e *Enemy) Left() {
	if e.X == 0 {
		return
	}
	e.X--
}

func (e *Enemy) Right() {
	if e.X == config.CountPointX-1 {
		return
	}
	e.X++
}

func (e *Enemy) Up() {
	if e.Y == 0 {
		return
	}
	e.Y--
}

func (e *Enemy) Down() {
	if e.Y == config.CountPointY-1 {

		return
	}
	e.Y++
}
