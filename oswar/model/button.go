package model

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Button struct {
	image     *ebiten.Image
	decision  image.Rectangle
	drawPoint image.Point
}

func NewButton(x, y int) (*Button, error) {
	img, err := ebiten.NewImage(160, 50, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{0x0, 0xf, 0x0, 0xff})

	d := img.Bounds()

	dp := image.Point{x, y}

	return &Button{
		image:     img,
		decision:  d.Add(dp),
		drawPoint: dp,
	}, nil
}

func (b *Button) ID() int {
	return 0
}

func (b *Button) Image() *ebiten.Image {
	return b.image
}

func (b *Button) Decision() *image.Rectangle {
	return &b.decision
}

func (b *Button) CenterX() int {
	return b.drawPoint.X + b.image.Bounds().Dx()/2
}

func (b *Button) CenterY() int {
	return b.drawPoint.Y + b.image.Bounds().Dy()/2
}

func (b *Button) Option() *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.drawPoint.X), float64(b.drawPoint.Y))

	return op
}

func (b *Button) IsPushed(x, y int) bool {
	p := image.Point{x, y}
	if p.X < b.decision.Min.X || b.decision.Max.X < p.X {
		return false
	}
	if p.Y < b.decision.Min.Y || b.decision.Max.Y < p.Y {
		return false
	}
	return true
}
