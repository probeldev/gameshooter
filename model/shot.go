package model

import (
	"github.com/probeldev/gameshooter/config"
)

type DirectionType int

const (
	DirectionTypeTop    DirectionType = 1
	DirectionTypeBottom DirectionType = 2
	DirectionTypeLeft   DirectionType = 3
	DirectionTypeRight  DirectionType = 4
)

type Shot struct {
	X int
	Y int

	Direction DirectionType
}

func NewShot(player Player) Shot {
	s := Shot{}

	var shotPositionX int = 0
	var shotPositionY int = 0

	pointSize := config.PointSize

	shotSize := config.GunSize
	playerStartX := player.X * pointSize
	playerStartY := player.Y * pointSize

	switch player.GunPosition {
	case GunPositionTypeTop:
		shotPositionX = playerStartX + (pointSize / 2)
		shotPositionY = playerStartY

		s.Direction = DirectionTypeTop

	case GunPositionTypeBottom:
		shotPositionX = playerStartX + (pointSize / 2)
		shotPositionY = playerStartY + pointSize - shotSize

		s.Direction = DirectionTypeBottom

	case GunPositionTypeLeft:
		shotPositionY = playerStartY + (pointSize / 2)
		shotPositionX = playerStartX

		s.Direction = DirectionTypeLeft

	case GunPositionTypeRight:
		shotPositionY = playerStartY + (pointSize / 2)
		shotPositionX = playerStartX + pointSize - shotSize

		s.Direction = DirectionTypeRight
	}

	s.X = shotPositionX
	s.Y = shotPositionY

	return s
}

func (s *Shot) Move() {
	switch s.Direction {
	case DirectionTypeTop:
		s.Y -= config.ShotSpeed
	case DirectionTypeBottom:
		s.Y += config.ShotSpeed
	case DirectionTypeLeft:
		s.X -= config.ShotSpeed
	case DirectionTypeRight:
		s.X += config.ShotSpeed
	}
}

func (s *Shot) IsNeedDelete() bool {

	if s.X < 0 {
		return true
	}

	if s.Y < 0 {
		return true
	}

	if s.X > config.WindowWidth {
		return true
	}
	if s.Y > config.WindowHeight {
		return true
	}

	return false
}
