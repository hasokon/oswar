package view

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Oswar is for the game OS War's View
type Oswar struct {
	gameLayer Layer
}

// NewOswar create Oswar instance
func NewOswar(screenWidth, screenHeight int) (*Oswar, error) {
	gl, err := NewGameLayer(screenWidth, screenHeight)
	if err != nil {
		return nil, err
	}

	return &Oswar{
		gameLayer: gl,
	}, nil
}

// GetUpdate make closer for ebiten.Run()
func (o *Oswar) GetUpdate() func(*ebiten.Image) error {
	return func(screen *ebiten.Image) error {
		if ebiten.IsRunningSlowly() {
			return nil
		}

		err := o.gameLayer.Update()
		if err != nil {
			return err
		}

		screen.DrawImage(o.gameLayer.Canvas(), nil)

		x, y := ebiten.CursorPosition()
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS : %0.2f\n(%d, %d)", ebiten.CurrentFPS(), x, y))
		return nil
	}
}
