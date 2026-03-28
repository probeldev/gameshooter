package config

import (
	"log"

	"golang.org/x/image/font"
)

const (
	PointSize    = 60
	GunSize      = 10
	WindowWidth  = 1280
	WindowHeight = 860
	MoveTime     = 15

	CountPointX = WindowWidth / PointSize
	CountPointY = WindowHeight / PointSize
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
	if PointSize%2 != 0 {
		log.Panic("PointSize может быть только четным")
	}

	if GunSize%2 != 0 {
		log.Panic("GunSize может быть только четным")
	}

	if PointSize <= GunSize {
		log.Panic("PointSize не может быть меньше или равным GunSize")
	}
}
