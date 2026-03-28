package screen

import (
	"image/color"
	"math"
	"strconv"

	"github.com/probeldev/gameshooter/config"
	"github.com/probeldev/gameshooter/model"
	"github.com/probeldev/gameshooter/scope"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type gameScreen struct {
	Player  model.Player
	Enemies []model.Enemy
	timer   int
	scope   *scope.Scope

	changeScreenFunc func(config.ScreenType)
}

func NewGameScreen(
	changeScreenFunc func(config.ScreenType),
	scope *scope.Scope,
) *gameScreen {
	gs := &gameScreen{}

	gs.Player = model.NewPlayer()

	gs.changeScreenFunc = changeScreenFunc

	gs.createEnemies()

	gs.scope = scope
	return gs
}

func (gs *gameScreen) createEnemies() {

	for range config.CountEnemies {
		gs.Enemies = append(gs.Enemies, model.NewEnemy())
	}

}

func (gs *gameScreen) Update() error {

	if gs.needsToMoveEnemy() {
		gs.moveEnemy()
	}

	if gs.isStopGame() {
		gs.changeScreenFunc(config.ScreenTypeGameOver)
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyH) {
		gs.Player.Left()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyL) {
		gs.Player.Right()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		gs.Player.Down()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyK) {
		gs.Player.Up()
	}

	gs.timer++
	return nil
}

func (gs *gameScreen) Draw(
	screenH *ebiten.Image,
) {
	// Обычный HUD с нормальным шрифтом

	screenH.Fill(color.RGBA{0x00, 0x33, 0x00, 0xFF})

	pointSize := float32(config.PointSize)

	gs.DrawPlayer(screenH)

	for _, enemy := range gs.Enemies {
		vector.FillRect(screenH, float32(enemy.X)*pointSize, float32(enemy.Y)*pointSize, pointSize, pointSize, color.RGBA{0xFF, 0x00, 0x00, 0xff}, false)
	}

	scoreText := "Score: " + strconv.Itoa(gs.scope.Value)
	text.Draw(screenH, scoreText, config.GameFont, 10, 30, color.White)
}

func (gs *gameScreen) DrawPlayer(
	screenH *ebiten.Image,
) {

	pointSize := float32(config.PointSize)

	playerStartX := float32(gs.Player.X) * pointSize
	playerStartY := float32(gs.Player.Y) * pointSize

	var gunPositionX float32 = 0
	var gunPositionY float32 = 0

	gunSize := float32(config.GunSize)

	switch gs.Player.GunPosition {
	case model.GunPositionTypeTop:
		gunPositionX = playerStartX + (pointSize/2 - config.GunSize/2)
		gunPositionY = playerStartY
	case model.GunPositionTypeBottom:
		gunPositionX = playerStartX + (pointSize/2 - config.GunSize/2)
		gunPositionY = playerStartY + pointSize - gunSize
	case model.GunPositionTypeLeft:
		gunPositionY = playerStartY + (pointSize/2 - config.GunSize/2)
		gunPositionX = playerStartX
	case model.GunPositionTypeRight:
		gunPositionY = playerStartY + (pointSize/2 - config.GunSize/2)
		gunPositionX = playerStartX + pointSize - gunSize
	}

	vector.FillRect(screenH, playerStartX, playerStartY, pointSize, pointSize, color.RGBA{0xFF, 0xFF, 0x00, 0x00}, false)

	vector.FillRect(screenH, gunPositionX, gunPositionY, gunSize, gunSize, color.RGBA{0x00, 0xFF, 0xFF, 0xff}, false)

}

func (gs *gameScreen) needsToMoveEnemy() bool {
	return gs.timer%config.MoveTime == 0
}

func (gs *gameScreen) isStopGame() bool {

	for _, enemy := range gs.Enemies {
		deltaX := gs.Player.X - enemy.X
		deltaY := gs.Player.Y - enemy.Y

		if deltaX == 0 && deltaY == 0 {
			return true
		}

	}

	return false
}

func (gs *gameScreen) moveEnemy() {
	// TODO:  Сделать что бы несколько врагов не занимало одну клеточку.
	for i := range gs.Enemies {
		deltaX := gs.Player.X - gs.Enemies[i].X
		deltaY := gs.Player.Y - gs.Enemies[i].Y

		if math.Abs(float64(deltaX)) > math.Abs(float64(deltaY)) {
			if deltaX > 0 {
				gs.Enemies[i].Right()
			} else if deltaX < 0 {
				gs.Enemies[i].Left()
			}
		} else {
			if deltaY > 0 {
				gs.Enemies[i].Down()
			} else if deltaY < 0 {
				gs.Enemies[i].Up()
			}
		}

	}

}
