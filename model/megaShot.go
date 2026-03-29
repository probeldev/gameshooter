package model

import (
	"log"

	"github.com/probeldev/gameshooter/config"
)

type DirectionMegaShotType int

const (
	DirectionMegaShotTypeTop    DirectionMegaShotType = 1
	DirectionMegaShotTypeBottom DirectionMegaShotType = 2
	DirectionMegaShotTypeLeft   DirectionMegaShotType = 3
	DirectionMegaShotTypeRight  DirectionMegaShotType = 4
)

type MegaShot struct {
	X int
	Y int

	Direction DirectionMegaShotType
}

func NewMegaShotFull(player Player) []MegaShot {
	full := []MegaShot{}

	full = append(full, NewMegaShot(player, DirectionMegaShotTypeTop))
	full = append(full, NewMegaShot(player, DirectionMegaShotTypeBottom))
	full = append(full, NewMegaShot(player, DirectionMegaShotTypeLeft))
	full = append(full, NewMegaShot(player, DirectionMegaShotTypeRight))

	return full
}

func NewMegaShot(
	player Player,
	direction DirectionMegaShotType,
) MegaShot {
	s := MegaShot{}

	var shotPositionX int = 0
	var shotPositionY int = 0

	s.Direction = direction

	pointSize := config.PointSize

	shotSize := config.GunSize
	playerStartX := player.X * pointSize
	playerStartY := player.Y * pointSize

	switch direction {
	case DirectionMegaShotTypeTop:
		shotPositionX = playerStartX + (pointSize / 2)
		shotPositionY = playerStartY

	case DirectionMegaShotTypeBottom:
		shotPositionX = playerStartX + (pointSize / 2)
		shotPositionY = playerStartY + pointSize - shotSize

	case DirectionMegaShotTypeLeft:
		shotPositionY = playerStartY + (pointSize / 2)
		shotPositionX = playerStartX

	case DirectionMegaShotTypeRight:
		shotPositionY = playerStartY + (pointSize / 2)
		shotPositionX = playerStartX + pointSize - shotSize

	}

	s.X = shotPositionX
	s.Y = shotPositionY

	return s
}

func (s *MegaShot) Move() {
	log.Println("MegaShot:Move 1", s)
	switch s.Direction {
	case DirectionMegaShotTypeTop:
		s.Y -= config.MegaShotSpeed
	case DirectionMegaShotTypeBottom:
		s.Y += config.MegaShotSpeed
	case DirectionMegaShotTypeLeft:
		s.X -= config.MegaShotSpeed
	case DirectionMegaShotTypeRight:
		s.X += config.MegaShotSpeed
	}
	log.Println("MegaShot:Move 2", s)
}

func (s *MegaShot) IsNeedDelete() bool {

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
