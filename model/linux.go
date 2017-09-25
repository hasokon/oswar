package model

import (
	"image"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/nfnt/resize"
)

type Linux struct {
	image    *ebiten.Image
	decision *image.Rectangle
}

func NewLinux() (*Linux, error) {
	_, img, err := ebitenutil.NewImageFromFile("resource/linux.png", ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}
	img = resize.Resize(100, 0, img, resize.Lanczos3)

	li, err := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	d := li.Bounds()

	return &Linux{
		image:    li,
		decision: &d,
	}, nil
}

func (l *Linux) Image() *ebiten.Image {
	return l.image
}

func (l *Linux) Dx() int {
	return l.image.Bounds().Dx()
}

func (l *Linux) Dy() int {
	return l.image.Bounds().Dy()
}
