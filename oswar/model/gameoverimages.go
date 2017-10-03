package model

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hasokon/oswar/oswar"
	"github.com/hasokon/oswar/oswar/controller"
)

type GameOverImages struct {
	BackgroundImage *ebiten.Image
	RestartButton   *Button
}

func NewGameOverImages() (*GameOverImages, error) {
	bgi, _, err := ebitenutil.NewImageFromFile("resource/gameover.png", ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	b, err := NewButton(420, 60)
	if err != nil {
		return nil, err
	}

	return &GameOverImages{
		BackgroundImage: bgi,
		RestartButton:   b,
	}, nil
}

func (goi *GameOverImages) MouseClicked(e controller.MouseEvent) error {
	if goi.RestartButton.IsPushed(e.X, e.Y) {
		oswar.SetState(oswar.IsPlaying)
	}
	return nil
}
