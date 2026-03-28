package config

import (
	"golang.org/x/image/font"
)

const (
	PointSize    = 60
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
