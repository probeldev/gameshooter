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
	if p.GunPosition != GunPositionTypeLeft {
		p.GunPosition = GunPositionTypeLeft
		return
	}

	if p.X == 0 {
		return
	}
	p.X--
}

func (p *Player) Right() {
	if p.GunPosition != GunPositionTypeRight {
		p.GunPosition = GunPositionTypeRight
		return
	}
	if p.X == config.CountPointX-1 {
		return
	}
	p.X++
}

func (p *Player) Up() {
	if p.GunPosition != GunPositionTypeTop {
		p.GunPosition = GunPositionTypeTop
		return
	}
	if p.Y == 0 {
		return
	}
	p.Y--

}

func (p *Player) Down() {
	if p.GunPosition != GunPositionTypeBottom {
		p.GunPosition = GunPositionTypeBottom
		return
	}
	if p.Y == config.CountPointY-1 {

		return
	}
	p.Y++
}
