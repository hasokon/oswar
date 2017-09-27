package model

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/nfnt/resize"

	"github.com/hajimehoshi/ebiten"
)

type GatesType int

const (
	GatesTypeVista GatesType = 0
	GatesTypeXP    GatesType = 1
	GatesType10    GatesType = 2
)

const (
	defaultLimmit = 15
)

const (
	SpeedLow    = 2
	SpeedMiddle = 3
	SpeedHigh   = 4
)

var (
	count      int
	gatesImage *ebiten.Image
	XPImage    *ebiten.Image
	VistaImage *ebiten.Image
	Win10Image *ebiten.Image
)

func GetLogoImage(path string) (*ebiten.Image, error) {
	_, img, err := ebitenutil.NewImageFromFile(path, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}
	img = resize.Resize(35, 0, img, resize.Lanczos3)

	li, err := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	return li, nil
}

// Gates is gates image class
type Gates struct {
	image     *ebiten.Image
	drawPoint image.Point //Left Up
	decision  image.Rectangle
	id        int
	speed     float64
	limmit    int
	killed    bool
	gatesType GatesType
	logoImage *ebiten.Image
}

// NewGates create gates instance
func NewGates(x, y int, t GatesType) (*Gates, error) {
	var err error
	if gatesImage == nil {
		gatesImage, _, err = ebitenutil.NewImageFromFile("resource/gates.png", ebiten.FilterNearest)
		if err != nil {
			return nil, err
		}
	}

	if XPImage == nil {
		VistaImage, err = GetLogoImage("resource/vista.png")
		if err != nil {
			return nil, err
		}

		XPImage, err = GetLogoImage("resource/xp.png")
		if err != nil {
			return nil, err
		}

		Win10Image, err = GetLogoImage("resource/win10.png")
		if err != nil {
			return nil, err
		}
	}

	r := gatesImage.Bounds()
	gatesDecisoin := &r

	count++

	var s float64
	var li *ebiten.Image
	switch t {
	case GatesTypeVista:
		s = SpeedLow
		li = VistaImage
	case GatesType10:
		s = SpeedMiddle
		li = Win10Image
	case GatesTypeXP:
		s = SpeedHigh
		li = XPImage
	default:
		s = SpeedLow
		li = VistaImage
	}

	// width, height := gatesImage.Size()
	// lwidth, _ := li.Size()
	// newImg, _ := ebiten.NewImage(width+lwidth, height, ebiten.FilterNearest)

	// newImg.DrawImage(gatesImage, nil)

	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(float64(width), 0)
	// newImg.DrawImage(li, op)

	return &Gates{
		image:     gatesImage,
		drawPoint: image.Point{x, y},
		decision:  gatesDecisoin.Add(image.Point{x, y}),
		id:        count,
		limmit:    defaultLimmit,
		killed:    false,
		speed:     s,
		gatesType: t,
		logoImage: li,
	}, nil
}

func (g *Gates) ID() int {
	return g.id
}

func (g *Gates) Decision() *image.Rectangle {
	return &g.decision
}

func (g *Gates) Kill() {
	g.killDefault()
}

func (g *Gates) killDefault() {
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
	}
	return nil
}

func (g *Gates) Option() *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.drawPoint.X),float64(g.drawPoint.Y))
	return op
}

// Translation updates coodinates
func (g *Gates) Translate(p image.Point) {
	g.drawPoint = g.drawPoint.Add(p)
	g.decision = g.decision.Add(p)
}

func (g *Gates) CenterX() int {
	return g.drawPoint.X + g.image.Bounds().Dx()/2
}

func (g *Gates) CenterY() int {
	return g.drawPoint.Y + g.image.Bounds().Dy()/2
}

func (g *Gates) Image() *ebiten.Image {
	width, height := g.image.Size()
	maskImg, _ := ebiten.NewImage(width, height, ebiten.FilterNearest)
	maskImg.Fill(color.Alpha{uint8(0xff * g.limmit / defaultLimmit)})

	op := &ebiten.DrawImageOptions{}
	op.CompositeMode = ebiten.CompositeModeSourceIn
	maskImg.DrawImage(g.image, op)

	return maskImg
}

func (g *Gates) MoveToPoint(dest image.Point) {
	dx := dest.X - g.CenterX()
	dy := dest.Y - g.CenterY()
	r := math.Sqrt(float64(dx*dx + dy*dy))

	// 中央で微妙に動いてずれるので、距離が単位速度未満で中央へ
	if r <= 0 {
		return
	}

	if r < g.speed {
		g.Translate(image.Point{dx, dy})
		return
	}

	sin := float64(dy) / r
	cos := float64(dx) / r

	x := int(g.speed * cos)
	y := int(g.speed * sin)

	g.Translate(image.Point{x, y})
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
