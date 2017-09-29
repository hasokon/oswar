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
func NewGameLayer(screenWidth, screenHeight int) (*GameLayer, error) {
	c, _ := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)

	gi, err := model.New(screenWidth, screenHeight)
	if err != nil {
		return nil, err
	}

	mm := controller.New()
	mm.AddMouseClickEventHandler(gi)

	return &GameLayer{
		canvas:       c,
		images:       gi,
		mouseManager: mm,
	}, nil
}

// Draw re-draw images
func (gl *GameLayer) Draw() {
	gl.canvas.Fill(color.White)

	gl.canvas.DrawImage(gl.images.BackGroundImage, nil)

	li := gl.images.LinuxImage
	gl.canvas.DrawImage(li.Image(), li.Option())

	for _, gates := range gl.images.GetGatesList() {
		gl.canvas.DrawImage(gates.Image(), gates.Option())
	}
}

func (gl *GameLayer) Update() error {
	err := gl.mouseManager.Update()
	if err != nil {
		return err
	}

	err = gl.images.Update()
	if err != nil {
		return err
	}

	gl.Draw()
	return nil
}

func (gl *GameLayer) Canvas() *ebiten.Image {
	return gl.canvas
}
