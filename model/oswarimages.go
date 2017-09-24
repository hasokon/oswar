package model

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hasokon/oswar/controller"
)

type OswarImages struct {
	CanvasImage *ebiten.Image
	GatesList   []*Gates
}

func New(canvasWidth, canvasHeight int) *OswarImages {
	c, _ := ebiten.NewImage(canvasWidth, canvasHeight, ebiten.FilterNearest)
	c.Fill(color.White)

	r, _ := ebiten.NewImage(20, 20, ebiten.FilterNearest)
	r.Fill(color.RGBA{0x0, 0x0, 0xff, 0xff})

	return &OswarImages{
		CanvasImage: c,
		GatesList:   make([]*Gates, 0),
	}
}

func (oi *OswarImages) DeleteGatesByID(id int) {
	newlist := make([]*Gates, 0)
	for _, gates := range oi.GatesList {
		if gates.ID != id {
			newlist = append(newlist, gates)
		}
	}
	oi.GatesList = newlist
}

func (oi *OswarImages) Update() error {
	for _, gates := range oi.GatesList {
		gates.Update()
		if gates.IsDead() {
			oi.DeleteGatesByID(gates.ID)
		}
	}
	return nil
}

func (oi *OswarImages) Do(e controller.MouseEvent) error {

	for i := len(oi.GatesList) - 1; i >= 0; i-- {
		gates := oi.GatesList[i]
		if gates.HitDecisionToPoint(image.Point{e.X, e.Y}) {
			gates.kill()
			return nil
		}
	}

	oi.GatesList = append(oi.GatesList, NewGates(e.X, e.Y))
	return nil
}
