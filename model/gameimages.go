package model

import (
	"image"
	"image/color"
	"math"
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
	GatesList          []*Gates
	LinuxImage         *Linux
	BackGroundImage    *ebiten.Image
	CanvasWidth        int
	CanvasHeight       int
	CanvasCenter       image.Point
	gatesGenerateCount int
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
		GatesList:          make([]*Gates, 0),
		LinuxImage:         li,
		BackGroundImage:    bgi,
		CanvasWidth:        canvasWidth,
		CanvasHeight:       canvasHeight,
		CanvasCenter:       image.Point{canvasWidth / 2, canvasHeight / 2},
		gatesGenerateCount: 0,
	}, nil
}

// DeleteGatesByID is to delete a Gates in GatesList by ID
func (gi *GameImages) DeleteGatesByID(id int) {
	newlist := make([]*Gates, 0)
	for _, gates := range gi.GatesList {
		if gates.ID != id {
			newlist = append(newlist, gates)
		}
	}
	gi.GatesList = newlist
}

// Update updates all images by time
func (gi *GameImages) Update() error {
	// Generate New Gates
	if gi.gatesGenerateCount == gatesGenerateTime {
		err := gi.NewGatesInCircle()
		if err != nil {
			return err
		}
		gi.gatesGenerateCount = 0
	} else {
		gi.gatesGenerateCount++
	}

	// Delete And Update Gates
	for _, gates := range gi.GatesList {
		if gates.IsDead() {
			gi.DeleteGatesByID(gates.ID)
		} else if hit, _ := gi.LinuxImage.HitDecisionToObject(gates); hit {
			gates.killSoon()
		} else {
			gates.UpdateImage()
			gates.MoveToPoint(gi.CanvasCenter)
		}
	}
	return nil
}

func (gi *GameImages) NewGatesInCircle() error {
	r := gi.CanvasWidth/2 + (rand.Intn(40) + 1)
	theta := rand.Intn(360) + 1

	theta_rad := math.Pi * float64(theta) / 180
	x := gi.CanvasCenter.X + int(float64(r)*math.Cos(theta_rad))
	y := gi.CanvasCenter.Y + int(float64(r)*math.Sin(theta_rad))

	ng, err := NewGates(x, y)
	if err != nil {
		return err
	}

	gi.GatesList = append(gi.GatesList, ng)
	return nil
}

// MouseClicked is execute by clicking mouse left button
func (gi *GameImages) MouseClicked(e controller.MouseEvent) error {
	for i := len(gi.GatesList) - 1; i >= 0; i-- {
		gates := gi.GatesList[i]
		if gates.HitDecisionToPoint(image.Point{e.X, e.Y}) {
			gates.killSoon()
			return nil
		}
	}
	return nil
}
