package view

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hasokon/oswar/oswar"
	"github.com/hasokon/oswar/oswar/controller"
	"github.com/hasokon/oswar/oswar/model"
)

type GameOverLayer struct {
	canvas       *ebiten.Image
	images       *model.GameOverImages
	mouseManager *controller.MouseEventListener
}

func NewGameOverLayer() (*GameOverLayer, error) {
	c, err := ebiten.NewImage(oswar.GetScreenWidth(), oswar.GetScreenHeight(), ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	imgs, err := model.NewGameOverImages()
	if err != nil {
		return nil, err
	}

	mm := controller.New()
	mm.AddMouseClickEventHandler(imgs)

	return &GameOverLayer{
		canvas:       c,
		images:       imgs,
		mouseManager: mm,
	}, nil
}

func (gol *GameOverLayer) Draw() {
	gol.canvas.Clear()

	gol.canvas.DrawImage(gol.images.BackgroundImage, nil)

	rb := gol.images.RestartButton
	gol.canvas.DrawImage(rb.Image(), rb.Option())
}

func (gol *GameOverLayer) Update() error {
	err := gol.mouseManager.Update()
	if err != nil {
		return err
	}

	gol.Draw()
	return nil
}

func (gol *GameOverLayer) Canvas() *ebiten.Image {
	return gol.canvas
}
