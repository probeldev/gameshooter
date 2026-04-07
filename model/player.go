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
	X           float64
	Y           float64
	GunPosition GunPositionType
}

func NewPlayer() Player {
	p := Player{}
	p.GunPosition = GunPositionTypeTop

	p.setSpawnPosition()

	return p
}

func (p *Player) setSpawnPosition() {
	p.X = float64(config.CountPointX / 2 * config.PlayerSize)
	p.Y = float64(config.CountPointY / 2 * config.PlayerSize)
}

func (p *Player) Left() {
	if p.GunPosition != GunPositionTypeLeft {
		p.GunPosition = GunPositionTypeLeft
		return
	}

	if p.X <= 0 {
		return
	}
	p.X -= config.PlayerSpeed
	if p.X < 0 {
		p.X = 0
	}
}

func (p *Player) Right() {
	if p.GunPosition != GunPositionTypeRight {
		p.GunPosition = GunPositionTypeRight
		return
	}
	if p.X >= config.WindowWidth-config.PlayerSize {
		return
	}
	p.X += config.PlayerSpeed
	if p.X > config.WindowWidth-config.PlayerSize {
		p.X = config.WindowWidth - config.PlayerSize
	}
}

func (p *Player) Up() {
	if p.GunPosition != GunPositionTypeTop {
		p.GunPosition = GunPositionTypeTop
		return
	}
	if p.Y <= 0 {
		return
	}
	p.Y -= config.PlayerSpeed
	if p.Y < 0 {
		p.Y = 0
	}
}

func (p *Player) Down() {
	if p.GunPosition != GunPositionTypeBottom {
		p.GunPosition = GunPositionTypeBottom
		return
	}
	if p.Y >= config.WindowHeight-config.PlayerSize {
		return
	}
	p.Y += config.PlayerSpeed
	if p.Y > config.WindowHeight-config.PlayerSize {
		p.Y = config.WindowHeight - config.PlayerSize
	}
}
