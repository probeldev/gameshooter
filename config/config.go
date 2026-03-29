package config

import (
	"log"

	"golang.org/x/image/font"
)

const (
	PlayerSize   = 40
	EnemySize    = 40
	GunSize      = 10
	ShotSize     = 10
	WindowWidth  = 1280
	WindowHeight = 860

	MoveEnemyPixel = 1

	ShotSpeed = 15
	ShotDelay = 10

	MegaShotSpeed = 15
	MegaShotDelay = 120

	CountPointX = WindowWidth / PlayerSize
	CountPointY = WindowHeight / PlayerSize
)

var GameFont font.Face

type ScreenType int

const (
	ScreenTypeStart    ScreenType = 1
	ScreenTypeGame     ScreenType = 2
	ScreenTypeGameOver ScreenType = 3
)

const CountEnemies = 10

func CheckConfig() {
	if PlayerSize%2 != 0 {
		log.Panic("PointSize может быть только четным")
	}

	if GunSize%2 != 0 {
		log.Panic("GunSize может быть только четным")
	}

	if PlayerSize <= GunSize {
		log.Panic("PointSize не может быть меньше или равным GunSize")
	}

	if ShotSize%2 != 0 {
		log.Panic("ShotSize может быть только четным")
	}

	if PlayerSize <= ShotSize {
		log.Panic("PointSize не может быть меньше или равным ShotSize")
	}
}
