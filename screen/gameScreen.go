package screen

import (
	"image/color"
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

	Shots []model.Shot

	MegaShot []model.MegaShot

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
		gs.Enemies = append(
			gs.Enemies,
			gs.CreateEnemy(),
		)
	}

}

func (gs *gameScreen) Update() error {

	gs.moveEnemy()

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

	// if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
	// 	gs.MegaShotRun()
	// }
	if gs.needsToMegaShot() {
		gs.MegaShotRun()
	}
	if gs.needsToShot() {
		gs.Shot()
	}

	gs.moveShots()
	gs.moveMegaShot()

	gs.killEnemyShot()
	gs.killEnemyMegaMegaShot()

	gs.timer++
	return nil
}

func (gs *gameScreen) moveShots() {

	for i := range gs.Shots {
		gs.Shots[i].Move()
	}

	shots := []model.Shot{}

	for i, shot := range gs.Shots {
		if gs.Shots[i].IsNeedDelete() {
			continue
		}

		shots = append(shots, shot)
	}

	gs.Shots = shots

}

func (gs *gameScreen) moveMegaShot() {

	for i := range gs.MegaShot {
		gs.MegaShot[i].Move()
	}

	shots := []model.MegaShot{}

	for i, shot := range gs.MegaShot {
		if gs.MegaShot[i].IsNeedDelete() {
			continue
		}

		shots = append(shots, shot)
	}

	gs.MegaShot = shots

}

func (gs *gameScreen) killEnemyMegaMegaShot() {

	deleteShotsIndex := []int{}
	deleteEnemiesIndex := []int{}

	for i, enemy := range gs.Enemies {
		for j, shot := range gs.MegaShot {
			if enemy.IsKillMegaShot(shot) {
				deleteShotsIndex = append(deleteShotsIndex, j)
				deleteEnemiesIndex = append(deleteEnemiesIndex, i)
				gs.scope.Value++
			}
		}
	}

	shots := []model.MegaShot{}
	enemies := []model.Enemy{}

	for i, shot := range gs.MegaShot {
		isNeedDelete := false
		for _, di := range deleteShotsIndex {
			if di == i {
				isNeedDelete = true
			}
		}

		if !isNeedDelete {
			shots = append(shots, shot)
		}
	}

	for i, enemy := range gs.Enemies {
		isNeedDelete := false
		for _, di := range deleteEnemiesIndex {
			if di == i {
				isNeedDelete = true
			}
		}

		if !isNeedDelete {
			enemies = append(enemies, enemy)
		}
	}

	gs.Enemies = enemies

	for range deleteEnemiesIndex {
		gs.Enemies = append(
			gs.Enemies,
			gs.CreateEnemy(),
		)
	}
	gs.MegaShot = shots

}

func (gs *gameScreen) killEnemyShot() {

	deleteShotsIndex := []int{}
	deleteEnemiesIndex := []int{}

	for i, enemy := range gs.Enemies {
		for j, shot := range gs.Shots {
			if enemy.IsKillShot(shot) {
				deleteShotsIndex = append(deleteShotsIndex, j)
				deleteEnemiesIndex = append(deleteEnemiesIndex, i)
				gs.scope.Value++
			}
		}
	}

	shots := []model.Shot{}
	enemies := []model.Enemy{}

	for i, shot := range gs.Shots {
		isNeedDelete := false
		for _, di := range deleteShotsIndex {
			if di == i {
				isNeedDelete = true
			}
		}

		if !isNeedDelete {
			shots = append(shots, shot)
		}
	}

	for i, enemy := range gs.Enemies {
		isNeedDelete := false
		for _, di := range deleteEnemiesIndex {
			if di == i {
				isNeedDelete = true
			}
		}

		if !isNeedDelete {
			enemies = append(enemies, enemy)
		}
	}

	gs.Enemies = enemies

	for range deleteEnemiesIndex {
		gs.Enemies = append(
			gs.Enemies,
			gs.CreateEnemy(),
		)
	}
	gs.Shots = shots

}

func (gs *gameScreen) CreateEnemy() model.Enemy {

	moved := model.NewEnemy(gs.Player)

	collided := false
	for _, other := range gs.Enemies {

		// AABB проверка пересечения (целочисленная)
		if moved.X < other.X+config.PlayerSize &&
			moved.X+config.PlayerSize > other.X &&
			moved.Y < other.Y+config.PlayerSize &&
			moved.Y+config.PlayerSize > other.Y {
			collided = true
			break
		}
	}

	// Если коллизия — пересоздаем
	if collided {
		return gs.CreateEnemy()
	}

	return moved
}

func (gs *gameScreen) Draw(
	screenH *ebiten.Image,
) {
	// Обычный HUD с нормальным шрифтом

	screenH.Fill(color.RGBA{0x00, 0x33, 0x00, 0xFF})

	gs.DrawEnemies(screenH)

	gs.DrawPlayer(screenH)
	gs.DrawShots(screenH)
	gs.DrawMegaShot(screenH)

	scoreText := "Score: " + strconv.Itoa(gs.scope.Value)
	text.Draw(screenH, scoreText, config.GameFont, 10, 30, color.White)
}

func (gs *gameScreen) Shot() {
	gs.Shots = append(gs.Shots, model.NewShot(gs.Player))
}

func (gs *gameScreen) MegaShotRun() {
	// TODO: Перезаписывает предыдущий MegaShot, стоит переделать.
	gs.MegaShot = model.NewMegaShotFull(gs.Player)
}

func (gs *gameScreen) DrawEnemies(
	screenH *ebiten.Image,
) {
	enemySize := float32(config.EnemySize)
	for _, enemy := range gs.Enemies {
		vector.FillRect(
			screenH,
			float32(enemy.X),
			float32(enemy.Y),
			enemySize,
			enemySize,
			color.RGBA{0xFF, 0x00, 0x00, 0xff},
			false,
		)
	}
}

func (gs *gameScreen) DrawShots(
	screenH *ebiten.Image,
) {
	for _, shot := range gs.Shots {
		vector.FillRect(
			screenH,
			float32(shot.X),
			float32(shot.Y),
			float32(config.ShotSize),
			float32(config.ShotSize),
			color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
			false,
		)
	}
}
func (gs *gameScreen) DrawMegaShot(
	screenH *ebiten.Image,
) {
	for _, shot := range gs.MegaShot {
		vector.FillRect(
			screenH,
			float32(shot.X),
			float32(shot.Y),
			float32(config.ShotSize),
			float32(config.ShotSize),
			color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
			false,
		)
	}
}

func (gs *gameScreen) DrawPlayer(
	screenH *ebiten.Image,
) {

	pointSize := float32(config.PlayerSize)

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

func (gs *gameScreen) needsToShot() bool {
	return gs.timer%config.ShotDelay == 0
}

func (gs *gameScreen) needsToMegaShot() bool {
	return gs.timer%config.MegaShotDelay == 0
}

func (gs *gameScreen) isStopGame() bool {
	// Вычисляем границы игрока один раз перед циклом
	playerStartX := gs.Player.X * config.PlayerSize
	playerEndX := gs.Player.X*config.PlayerSize + config.PlayerSize
	playerStartY := gs.Player.Y * config.PlayerSize
	playerEndY := gs.Player.Y*config.PlayerSize + config.PlayerSize

	for _, enemy := range gs.Enemies {
		enemyStartX := enemy.X
		enemyEndX := enemy.X + config.EnemySize
		enemyStartY := enemy.Y
		enemyEndY := enemy.Y + config.EnemySize

		// Проверка пересечения двух прямоугольников (AABB collision)
		// Прямоугольники НЕ пересекаются, если один из них:
		// - полностью слева, справа, сверху или снизу от другого
		// Значит, пересекаются, если это условие ложно
		if playerStartX < enemyEndX &&
			playerEndX > enemyStartX &&
			playerStartY < enemyEndY &&
			playerEndY > enemyStartY {
			return true
		}
	}

	return false
}

func (gs *gameScreen) moveEnemy() {
	for i := range gs.Enemies {
		deltaX := gs.Player.X*config.PlayerSize - gs.Enemies[i].X
		deltaY := gs.Player.Y*config.PlayerSize - gs.Enemies[i].Y

		// Двигаем по оси с наибольшей разницей
		if abs(deltaX) > abs(deltaY) {
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

		// Проверяем коллизию с другими врагами
		moved := &gs.Enemies[i]
		collided := false

		for j, other := range gs.Enemies {
			if i == j {
				continue
			}

			// AABB проверка пересечения (целочисленная)
			if moved.X < other.X+config.PlayerSize &&
				moved.X+config.PlayerSize > other.X &&
				moved.Y < other.Y+config.PlayerSize &&
				moved.Y+config.PlayerSize > other.Y {
				collided = true
				break
			}
		}

		// Если коллизия — отменяем движение
		if collided {
			gs.Enemies[i].BackMove()
		}
	}
}

// Вспомогательная функция для abs(int), чтобы не тащить math.Abs для float64
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
