package model

import (
	"math/rand"

	"github.com/probeldev/gameshooter/config"
)

type Enemy struct {
	X     int
	Y     int
	PrevX int
	PrevY int
}

func NewEnemy(
	player Player,
) Enemy {
	e := Enemy{}
	e.setSpawnPosition(player)

	return e
}

func (e *Enemy) setSpawnPosition(
	player Player,
) {
	positions := []string{}

	// Делаем, что бы враги не спавнились прямо около игрока.
	if player.X > 5 {
		positions = append(positions, "left")
	}

	if player.X < config.CountPointX-1-5 {
		positions = append(positions, "right")
	}

	if player.Y > 5 {
		positions = append(positions, "top")
	}

	if player.Y < config.CountPointY-1-5 {
		positions = append(positions, "bottom")
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
	e.savePrev()
	if e.X == 0 {
		return
	}
	e.X--
}

func (e *Enemy) Right() {
	e.savePrev()
	if e.X == config.CountPointX-1 {
		return
	}
	e.X++
}

func (e *Enemy) Up() {
	e.savePrev()
	if e.Y == 0 {
		return
	}
	e.Y--
}

func (e *Enemy) Down() {
	e.savePrev()
	if e.Y == config.CountPointY-1 {

		return
	}
	e.Y++
}

func (e *Enemy) savePrev() {
	e.PrevX = e.X
	e.PrevY = e.Y
}

func (e *Enemy) BackMove() {
	// Возвращается на предыдущую позицию

	e.X = e.PrevX
	e.Y = e.PrevY
}

func (e *Enemy) IsKillShot(shot Shot) bool {
	enemyStartX := e.X * config.PointSize
	enemyEndX := (e.X + 1) * config.PointSize

	enemyStartY := e.Y * config.PointSize
	enemyEndY := (e.Y + 1) * config.PointSize

	if shot.X > enemyStartX && shot.X < enemyEndX &&
		shot.Y > enemyStartY && shot.Y < enemyEndY {
		return true
	}

	return false
}

func (e *Enemy) IsKillMegaShot(shot MegaShot) bool {
	enemyStartX := e.X * config.PointSize
	enemyEndX := (e.X + 1) * config.PointSize

	enemyStartY := e.Y * config.PointSize
	enemyEndY := (e.Y + 1) * config.PointSize

	if shot.X > enemyStartX && shot.X < enemyEndX &&
		shot.Y > enemyStartY && shot.Y < enemyEndY {
		return true
	}

	return false
}
