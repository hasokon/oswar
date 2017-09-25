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
func NewOswar(screenWidth, screenHeight int) *Oswar {
	gl := NewGameLayer(screenWidth, screenHeight)

	return &Oswar{
		gameLayer: gl,
	}
}

// GetUpdate make closer for ebiten.Run()
func (o *Oswar) GetUpdate() func(*ebiten.Image) error {
	return func(screen *ebiten.Image) error {
		o.gameLayer.Update()
		screen.DrawImage(o.gameLayer.Canvas(), nil)

		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS : %0.2f", ebiten.CurrentFPS()))
		return nil
	}
}
