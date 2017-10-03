package view

import (
	"fmt"

	"github.com/hasokon/oswar/oswar"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Oswar is for the game OS War's View
type Oswar struct {
	gameLayer Layer
	gameOverLayer Layer
}

// NewOswar create Oswar instance
func NewOswar(screenWidth, screenHeight int) (*Oswar, error) {
	gl, err := NewGameLayer(screenWidth, screenHeight)
	if err != nil {
		return nil, err
	}

	gol, err := NewGameOverLayer()
	if err != nil {
		return nil, err
	}

	return &Oswar{
		gameLayer: gl,
		gameOverLayer: gol,
	}, nil
}

func (o *Oswar) Update(screen *ebiten.Image) error {

	var cl Layer
	switch (oswar.GetState()) {
	case oswar.IsPlaying:
		cl = o.gameLayer
	case oswar.GameOver:
		cl = o.gameOverLayer
	default:
		cl = o.gameLayer
	}

	err := cl.Update()
	if err != nil {
		return err
	}
	screen.DrawImage(cl.Canvas(), nil)
	return nil
}

func Update() func(*ebiten.Image) error {
	o, _ := NewOswar(oswar.GetScreenWidth(), oswar.GetScreenHeight())

	return func(screen *ebiten.Image) error {
		if ebiten.IsRunningSlowly() {
			return nil
		}

		o.Update(screen)

		x, y := ebiten.CursorPosition()
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS : %0.2f\n(%d, %d)", ebiten.CurrentFPS(), x, y))
		return nil
	}
}