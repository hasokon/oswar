package view

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hasokon/oswar/controller"
	"github.com/hasokon/oswar/model"
)

// Oswar is for the game OS War's View
type Oswar struct {
	images       *model.OswarImages
	mouseManager *controller.MouseEventListener
}

// New create Oswar instance
func New(screenWidth, screenHeight int) *Oswar {
	oi := model.New(screenWidth, screenHeight)
	mm := controller.New()
	mm.AddMouseClickEventHandler(oi)

	return &Oswar{
		images:       oi,
		mouseManager: mm,
	}
}

// Update re-draw images
func (o *Oswar) Update(screen *ebiten.Image) error {
	screen.DrawImage(o.images.CanvasImage, nil)

	for _, gates := range o.images.GatesList {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(gates.DrawPoint.X), float64(gates.DrawPoint.Y))
		op.ColorM.RotateHue(float64(gates.ID))

		screen.DrawImage(gates.Image, op)
	}
	return nil
}

// GetUpdate make closer for ebiten.Run()
func (o *Oswar) GetUpdate() func(*ebiten.Image) error {
	return func(screen *ebiten.Image) error {
		o.mouseManager.Update()
		o.images.Update()
		o.Update(screen)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS : %0.2f", ebiten.CurrentFPS()))
		return nil
	}
}
