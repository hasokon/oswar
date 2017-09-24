package model

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var (
	gatesImage    *ebiten.Image
	gatesDecisoin *image.Rectangle
	count         int
)

type Gates struct {
	Image     *ebiten.Image
	DrawPoint *image.Point //Left Up
	decision  *image.Rectangle
	ID        int
}

func NewGates(x, y int) Gates {
	if gatesImage == nil {
		gatesImage, _ = ebiten.NewImage(40, 40, ebiten.FilterNearest)
		gatesImage.Fill(color.RGBA{0x0, 0x0, 0xf0, 0xff})

		r := image.Rect(0, 0, 40, 40)
		gatesDecisoin = &r

		count = 0
	}

	count++
	dp := image.Point{x, y}
	d := gatesDecisoin.Add(dp)

	return Gates{
		Image:     gatesImage,
		DrawPoint: &dp,
		decision:  &d,
		ID:        count,
	}
}

func (g *Gates) Translation(p image.Point) {
	g.DrawPoint.Add(p)
	g.decision.Add(p)
}

func (g *Gates) HitDecisionToPoint(p image.Point) bool {
	if p.X < g.decision.Min.X || g.decision.Max.X < p.X {
		return false
	}
	if p.Y < g.decision.Min.Y || g.decision.Max.Y < p.Y {
		return false
	}
	return true
}
