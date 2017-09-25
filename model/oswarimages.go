package model

import (
	"image"
	"image/color"
	"math/rand"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hasokon/oswar/controller"
)

const (
	gatesGenerateTime int = 60
)

// OswarImages have all image for this game
type OswarImages struct {
	CanvasImage *ebiten.Image
	GatesList   []*Gates
	ScreenWidth  int
	ScreenHeight int
	ScreenCenter image.Point
	gatesGenerateCount int
}

// New create OswarImages instance
func New(canvasWidth, canvasHeight int) *OswarImages {
	c, _ := ebiten.NewImage(canvasWidth, canvasHeight, ebiten.FilterNearest)
	c.Fill(color.White)

	rand.Seed(time.Now().UnixNano())

	return &OswarImages{
		CanvasImage: 		c,
		GatesList:   		make([]*Gates, 0),
		ScreenWidth: 		canvasWidth,
		ScreenHeight: 		canvasHeight,
		ScreenCenter:		image.Point{canvasWidth/2, canvasHeight/2},
		gatesGenerateCount:	0,
	}
}

// DeleteGatesByID is to delete a Gates in GatesList by ID
func (oi *OswarImages) DeleteGatesByID(id int) {
	newlist := make([]*Gates, 0)
	for _, gates := range oi.GatesList {
		if gates.ID != id {
			newlist = append(newlist, gates)
		}
	}
	oi.GatesList = newlist
}

// Update updates all images by time
func (oi *OswarImages) Update() error {
	// Generate New Gates
	if oi.gatesGenerateCount == gatesGenerateTime {
		oi.NewGatesInCircle()
		oi.gatesGenerateCount = 0
	} else {
		oi.gatesGenerateCount++
	}

	// Delete And Update Gates
	for _, gates := range oi.GatesList {
		if gates.IsDead() {
			oi.DeleteGatesByID(gates.ID)
		} else {
			gates.UpdateImage()
			gates.MoveToPoint(oi.ScreenCenter)
		}
	}
	return nil
}

func (oi *OswarImages) NewGatesInCircle() {
	r := oi.ScreenWidth/2 + (rand.Intn(40) + 1)
	theta := rand.Intn(360) + 1

	theta_rad := math.Pi*float64(theta)/180
	x := oi.ScreenCenter.X + int(float64(r) * math.Cos(theta_rad))
	y := oi.ScreenCenter.Y + int(float64(r) * math.Sin(theta_rad))

	oi.GatesList = append(oi.GatesList, NewGates(x,y))
}

// MouseClicked is execute by clicking mouse left button
func (oi *OswarImages) MouseClicked(e controller.MouseEvent) error {
	for i := len(oi.GatesList) - 1; i >= 0; i-- {
		gates := oi.GatesList[i]
		if gates.HitDecisionToPoint(image.Point{e.X, e.Y}) {
			gates.killSoon()
			return nil
		}
	}
	return nil
}
