package model

import (
	"image"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/nfnt/resize"
)

type Linux struct {
	image     *ebiten.Image
	decision  *image.Rectangle
	DrawPoint *image.Point
}

func NewLinux(x, y int) (*Linux, error) {
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

	drawX := x - li.Bounds().Dx()/2
	drawY := y - li.Bounds().Dy()/2
	dp := image.Point{drawX, drawY}

	d = d.Add(dp)

	return &Linux{
		image:     li,
		decision:  &d,
		DrawPoint: &dp,
	}, nil
}

func (l *Linux) Image() *ebiten.Image {
	return l.image
}

func (l *Linux) Decision() *image.Rectangle {
	return l.decision
}

func (l *Linux) Dx() int {
	return l.image.Bounds().Dx()
}

func (l *Linux) Dy() int {
	return l.image.Bounds().Dy()
}

func (l *Linux) CenterX() int {
	return l.DrawPoint.X + l.image.Bounds().Dx()
}

func (l *Linux) CenterY() int {
	return l.DrawPoint.Y + l.image.Bounds().Dy()
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}

func (l *Linux) HitDecisionToObject(o Object) (bool, *image.Point) {
	destDecision := o.Decision()
	srcDecision := l.Decision()

	dhdx := destDecision.Dx() / 2
	dhdy := destDecision.Dy() / 2
	shdx := srcDecision.Dx() / 2
	shdy := srcDecision.Dy() / 2

	destCenter := destDecision.Min.Add(image.Point{dhdx, dhdy})
	srcCenter := srcDecision.Min.Add(image.Point{shdx, shdy})

	dX := abs(destCenter.X - srcCenter.X)
	dY := abs(destCenter.Y - srcCenter.Y)

	mX := dhdx + shdx
	mY := dhdy + shdy

	if dX >= mX {
		return false, nil
	}

	if dY >= mY {
		return false, nil
	}

	nX := (destCenter.X + srcCenter.X) / 2
	nY := (destCenter.Y + srcCenter.Y) / 2

	return true, &image.Point{nX, nY}
}
