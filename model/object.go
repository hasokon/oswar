package model

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type Object interface {
	Image() *ebiten.Image
	Decision() *image.Rectangle
	CenterX() int
	CenterY() int
	Option() *ebiten.DrawImageOptions
	ID() int
}
