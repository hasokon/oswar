package model

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	defaultLimmit = 30
)

var (
	count int
)

// Gates is gates image class
type Gates struct {
	Image     *ebiten.Image
	DrawPoint image.Point //Left Up
	decision  image.Rectangle
	ID        int
	limmit    int
	killed    bool
}

// NewGates create gates instance
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
		limmit:    defaultLimmit,
		killed:    false,
	}
}

func (g *Gates) kill() {
	g.killed = true
}

func (g *Gates) killSoon() {
	g.killed = true
	g.limmit = 0
}

// IsDead express which this gates is dead or alive
func (g *Gates) IsDead() bool {
	return g.limmit <= 0
}

// Update updates gates by time
func (g *Gates) Update() error {
	if g.killed {
		g.limmit--
		g.Image.Fill(color.RGBA{0x0, 0x0, 0xf0, uint8(0xff * g.limmit / defaultLimmit)})
	}

	g.Translation(image.Point{1, 1})

	return nil
}

// Translation updates coodinates
func (g *Gates) Translation(p image.Point) {
	g.DrawPoint = g.DrawPoint.Add(p)
	g.decision = g.decision.Add(p)
}

// HitDecisionToPoint is decide wether a point is in this
func (g *Gates) HitDecisionToPoint(p image.Point) bool {
	if p.X < g.decision.Min.X || g.decision.Max.X < p.X {
		return false
	}
	if p.Y < g.decision.Min.Y || g.decision.Max.Y < p.Y {
		return false
	}
	return true
}
