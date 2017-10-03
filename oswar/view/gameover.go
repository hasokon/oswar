package view

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hasokon/oswar/oswar/controller"
	"github.com/hasokon/oswar/oswar/model"
)

type GameOverLayer struct {
	images       *model.GameOverImages
	mouseManager *controller.MouseEventListener
}

func NewGameOverLayer() (*GameOverLayer, error) {
	imgs, err := model.NewGameOverImages()
	if err != nil {
		return nil, err
	}

	mm := controller.New()
	mm.AddMouseClickEventHandler(imgs)

	return &GameOverLayer{
		images:       imgs,
		mouseManager: mm,
	}, nil
}

func (gol *GameOverLayer) Update() error {
	err := gol.mouseManager.Update()
	if err != nil {
		return err
	}

	return nil
}

func (gol *GameOverLayer) Canvas() *ebiten.Image {
	return gol.images.BackgroundImage
}
