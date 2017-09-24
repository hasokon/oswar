package model

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	default_limmit = 30
)

var (
	count int = 0
)

type Gates struct {
	Image     *ebiten.Image
	DrawPoint image.Point //Left Up
	decision  image.Rectangle
	ID        int
	limmit    int
	killed    bool
}

func NewGates(x, y int) *Gates {
	gatesImage, _ := ebiten.NewImage(40, 40, ebiten.FilterNearest)
	gatesImage.Fill(color.RGBA{0x0, 0x0, 0xf0, 0xff})
	r := image.Rect(0, 0, 40, 40)
	gatesDecisoin := &r

	count++

	return &Gates{
		Image:     gatesImage,
		DrawPoint: image.Point{x, y},
		decision:  gatesDecisoin.Add(image.Point{x, y}),
		ID:        count,
		limmit:    default_limmit,
		killed:    false,
	}
}

func (g *Gates) kill() {
	g.killed = true
}

func (g *Gates) IsDead() bool {
	return g.limmit <= 0
}

func (g *Gates) Update() error {
	if g.killed {
		g.limmit--
		g.Image.Fill(color.RGBA{0x0, 0x0, 0xf0, uint8(0xff * g.limmit / default_limmit)})
	}
	return nil
}

func (g *Gates) Translation(p image.Point) {
	g.DrawPoint = g.DrawPoint.Add(p)
	g.decision = g.decision.Add(p)
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
