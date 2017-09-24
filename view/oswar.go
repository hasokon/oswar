package view

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hasokon/oswar/controller"
	"github.com/hasokon/oswar/model"
)

type Oswar struct {
	screenWidth  int
	screenHeight int
	images       *model.OswarImages
	mouseManager *controller.MouseEventListener
}

func New(screenWidth, screenHeight int) *Oswar {
	oi := model.New(screenWidth, screenHeight)
	mm := controller.New()
	mm.AddMouseClickEventHandler(oi)

	return &Oswar{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		images:       oi,
		mouseManager: mm,
	}
}

func (o *Oswar) draw(screen *ebiten.Image) error {
	screen.DrawImage(o.images.CanvasImage, nil)

	for _, gates := range o.images.GatesList {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(gates.DrawPoint.X), float64(gates.DrawPoint.Y))
		op.ColorM.RotateHue(float64(gates.ID))

		screen.DrawImage(gates.Image, op)
	}
	return nil
}

func (o *Oswar) GetUpdate() func(*ebiten.Image) error {
	return func(screen *ebiten.Image) error {
		o.mouseManager.Update()
		o.draw(screen)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS : %0.2f", ebiten.CurrentFPS()))
		return nil
	}
}
