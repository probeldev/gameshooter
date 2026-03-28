package model

import "github.com/probeldev/gameshooter/config"

type GunPositionType int

const (
	GunPositionTypeTop    = 1
	GunPositionTypeBottom = 2
	GunPositionTypeLeft   = 3
	GunPositionTypeRight  = 4
)

type Player struct {
	X           int
	Y           int
	GunPosition GunPositionType
}

func NewPlayer() Player {
	p := Player{}
	p.GunPosition = GunPositionTypeTop

	p.setSpawnPosition()

	return p
}

func (p *Player) setSpawnPosition() {
	p.X = config.CountPointX / 2
	p.Y = config.CountPointY / 2
}

func (p *Player) Left() {
	if p.X == 0 {
		return
	}
	p.X--

	p.GunPosition = GunPositionTypeLeft
}

func (p *Player) Right() {
	if p.X == config.CountPointX-1 {
		return
	}
	p.X++

	p.GunPosition = GunPositionTypeRight
}

func (p *Player) Up() {
	if p.Y == 0 {
		return
	}
	p.Y--

	p.GunPosition = GunPositionTypeTop
}

func (p *Player) Down() {
	if p.Y == config.CountPointY-1 {

		return
	}
	p.Y++

	p.GunPosition = GunPositionTypeBottom
}
