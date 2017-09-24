package view

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hasokon/oswar/controller"
	"github.com/hasokon/oswar/model"
)

type Oswar struct {
	screenWidth  int
	screenHeight int
	images       *model.OswarImages
	cont         *controller.OswarController
}

func New(screenWidth, screenHeight int) *Oswar {
	oi := model.New(screenWidth, screenHeight)
	oc := controller.New()
	oc.AddMouseClickEventHandler(oi)

	return &Oswar{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		images:       oi,
		cont:         oc,
	}
}

func (o *Oswar) draw(screen *ebiten.Image) error {
	screen.DrawImage(o.images.CanvasImage, nil)
	return nil
}

func (o *Oswar) GetUpdate() func(*ebiten.Image) error {
	return func(screen *ebiten.Image) error {
		o.cont.Update()
		o.draw(screen)
		return nil
	}
}
