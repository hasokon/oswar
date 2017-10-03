package view

import (
	"github.com/hajimehoshi/ebiten"
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

func (o *Oswar) Update(screen *ebiten.Image) error {
	err := o.gameLayer.Update()
	if err != nil {
		return err
	}

	screen.DrawImage(o.gameLayer.Canvas(), nil)
	return nil
}
