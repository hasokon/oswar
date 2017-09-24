package model

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hasokon/oswar/controller"
)

type OswarImages struct {
	CanvasImage *ebiten.Image
	rect        *ebiten.Image
}

func New(canvasWidth, canvasHeight int) *OswarImages {
	c, _ := ebiten.NewImage(canvasWidth, canvasHeight, ebiten.FilterNearest)
	c.Fill(color.White)

	r, _ := ebiten.NewImage(20, 20, ebiten.FilterNearest)
	r.Fill(color.RGBA{0x0, 0x0, 0xff, 0xff})

	return &OswarImages{
		CanvasImage: c,
		rect:        r,
	}
}

func (oi *OswarImages) Do(e controller.MouseEvent) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.X), float64(e.Y))

	oi.CanvasImage.DrawImage(oi.rect, op)

	return nil
}
