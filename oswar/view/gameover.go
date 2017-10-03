package view

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type GameOverLayer struct {
	backgroundImage	*ebiten.Image
}

func NewGameOverLayer() (*GameOverLayer, error) {
	bgi, _, err := ebitenutil.NewImageFromFile("resource/gameover.png", ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	return &GameOverLayer {
		backgroundImage: bgi,
	}, nil
}

func (gol *GameOverLayer) Update() error {
	return nil
}

func (gol *GameOverLayer) Canvas() *ebiten.Image {
	return gol.backgroundImage
}