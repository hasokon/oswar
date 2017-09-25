package model

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

const (
	defaultLimmit = 30
	SpeedDefault = 2
	SpeedMiddle = 3
	SpeedHigh = 4
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
	Speed	  float64
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
		Speed:     SpeedDefault,
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

// UpdateImage updates gates's image by time
func (g *Gates) UpdateImage() error {
	if g.killed {
		g.limmit--
		g.Image.Fill(color.RGBA{0x0, 0x0, 0xf0, uint8(0xff * g.limmit / defaultLimmit)})
	}
	return nil
}

// Translation updates coodinates
func (g *Gates) Translation(p image.Point) {
	g.DrawPoint = g.DrawPoint.Add(p)
	g.decision = g.decision.Add(p)
}

func (g *Gates) CenterX() int {
	return g.DrawPoint.X + g.Image.Bounds().Dx()
}

func (g *Gates) CenterY() int {
	return g.DrawPoint.Y + g.Image.Bounds().Dy()
}

func (g *Gates) MoveToPoint(dest image.Point) {
	dx := dest.X - g.CenterX()
	dy := dest.Y - g.CenterY()
	r := math.Sqrt(float64(dx*dx + dy*dy))

	// 中央で微妙に動いてずれるので、距離が単位速度未満で中央へ
	if r <= 0 {
		return
	}
	
	if r < g.Speed {
		g.Translation(image.Point{dx,dy})
		return
	}

	sin := float64(dy) / r
	cos := float64(dx) / r

	x := int(g.Speed * cos)
	y := int(g.Speed * sin)

	g.Translation(image.Point{x,y})
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
