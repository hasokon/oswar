package model

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hasokon/oswar/controller"
)

const (
	gatesGenerateTime int = 60
)

// GameImages have all image for this game
type GameImages struct {
	CanvasImage        *ebiten.Image
	gatesList          *GatesList
	LinuxImage         *Linux
	BackGroundImage    *ebiten.Image
	CanvasWidth        int
	CanvasHeight       int
	CanvasCenter       image.Point
	gatesGenerateCount int
}

func (gi *GameImages) GetGatesList() []*Gates {
	return gi.gatesList.GetList()
}

// New create GameImages instance
func New(canvasWidth, canvasHeight int) (*GameImages, error) {
	c, _ := ebiten.NewImage(canvasWidth, canvasHeight, ebiten.FilterNearest)
	c.Fill(color.White)

	li, err := NewLinux(canvasWidth/2, canvasHeight/2)
	if err != nil {
		return nil, err
	}

	bgi, _, err := ebitenutil.NewImageFromFile("resource/gamebg.png", ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())

	return &GameImages{
		CanvasImage:        c,
		gatesList:          NewGatesList(),
		LinuxImage:         li,
		BackGroundImage:    bgi,
		CanvasWidth:        canvasWidth,
		CanvasHeight:       canvasHeight,
		CanvasCenter:       image.Point{canvasWidth / 2, canvasHeight / 2},
		gatesGenerateCount: 0,
	}, nil
}

// Update updates all images by time
func (gi *GameImages) Update() error {
	// Generate New Gates
	if gi.gatesGenerateCount == gatesGenerateTime {
		err := gi.gatesList.NewGatesInCircle(gi.CanvasWidth, gi.CanvasHeight, gi.CanvasCenter)
		if err != nil {
			return err
		}
		gi.gatesGenerateCount = 0
	} else {
		gi.gatesGenerateCount++
	}

	// Delete And Update Gates
	for _, gates := range gi.gatesList.GetList() {
		if gates.IsDead() {
			gi.gatesList.DeleteGatesByID(gates.ID())
		} else if hit, _ := gi.LinuxImage.HitDecisionToObject(gates); hit {
			gates.Kill()
		}
		gates.UpdateImage()
		gates.MoveToPoint(gi.CanvasCenter)
	}
	return nil
}

// MouseClicked is execute by clicking mouse left button
func (gi *GameImages) MouseClicked(e controller.MouseEvent) error {
	gl := gi.gatesList.GetList()
	for i := len(gl) - 1; i >= 0; i-- {
		gates := gl[i]
		if gates.HitDecisionToPoint(image.Point{e.X, e.Y}) {
			gates.Kill()
			return nil
		}
	}
	return nil
}
