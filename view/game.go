package view

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hasokon/oswar/controller"
	"github.com/hasokon/oswar/model"
)

type GameLayer struct {
	canvas       *ebiten.Image
	images       *model.GameImages
	mouseManager *controller.MouseEventListener
}

// NewGameLayer create GameImage instance
func NewGameLayer(screenWidth, screenHeight int) *GameLayer {
	c, _ := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)

	gi := model.New(screenWidth, screenHeight)
	mm := controller.New()
	mm.AddMouseClickEventHandler(gi)

	return &GameLayer{
		canvas:       c,
		images:       gi,
		mouseManager: mm,
	}
}

// Draw re-draw images
func (gl *GameLayer) Draw() {
	gl.canvas.Fill(color.White)

	for _, gates := range gl.images.GatesList {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(gates.DrawPoint.X), float64(gates.DrawPoint.Y))
		op.ColorM.RotateHue(float64(gates.ID))

		gl.canvas.DrawImage(gates.Image, op)
	}
}

func (gl *GameLayer) Update() error {
	gl.mouseManager.Update()
	gl.images.Update()
	gl.Draw()
	return nil
}

func (gl *GameLayer) Canvas() *ebiten.Image {
	return gl.canvas
}
