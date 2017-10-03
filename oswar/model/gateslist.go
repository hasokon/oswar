package model

import (
	"image"
	"math"
	"math/rand"
)

const (
	defaultSliceLength = 32
)

type GatesList struct {
	usingGatesList []*Gates
	xpGatesList    []*Gates
	vistaGatesList []*Gates
	win10GatesList []*Gates
}

func NewGatesList() *GatesList {
	return &GatesList{
		usingGatesList: make([]*Gates, 0, defaultSliceLength),
		xpGatesList:    make([]*Gates, 0, defaultSliceLength),
		vistaGatesList: make([]*Gates, 0, defaultSliceLength),
		win10GatesList: make([]*Gates, 0, defaultSliceLength),
	}
}

func (gl *GatesList) Reset() {
	for _, gates := range gl.usingGatesList {
		switch gates.Type() {
		case GatesTypeXP:
			gl.xpGatesList = append(gl.xpGatesList, gates)
		case GatesTypeVista:
			gl.vistaGatesList = append(gl.vistaGatesList, gates)
		case GatesType10:
			gl.win10GatesList = append(gl.win10GatesList, gates)
		}
	}

	gl.usingGatesList = make([]*Gates, 0, defaultSliceLength)
}

func (gl *GatesList) GetList() []*Gates {
	return gl.usingGatesList
}

// DeleteGatesByID is to delete a Gates in GatesList by ID
func (gl *GatesList) DeleteGatesByID(id int) {
	newlist := make([]*Gates, 0, defaultSliceLength)
	for _, gates := range gl.usingGatesList {
		if gates.ID() != id {
			newlist = append(newlist, gates)
		} else {
			switch gates.Type() {
			case GatesTypeXP:
				gl.xpGatesList = append(gl.xpGatesList, gates)
			case GatesTypeVista:
				gl.vistaGatesList = append(gl.vistaGatesList, gates)
			case GatesType10:
				gl.win10GatesList = append(gl.win10GatesList, gates)
			}
		}
	}
	gl.usingGatesList = newlist
}

func (gl *GatesList) NewGatesInCircle(width, height int, center image.Point) error {
	r := width/2 + (rand.Intn(40) + 1)
	theta := rand.Intn(360) + 1

	theta_rad := math.Pi * float64(theta) / 180
	x := center.X + int(float64(r)*math.Cos(theta_rad))
	y := center.Y + int(float64(r)*math.Sin(theta_rad))

	p := rand.Intn(100) + 1

	ng, err := gl.NewGates(x, y, p)
	if err != nil {
		return err
	}

	gl.usingGatesList = append(gl.usingGatesList, ng)

	return nil
}

func (gl *GatesList) NewGates(x, y, rand int) (*Gates, error) {
	var gtype GatesType
	switch {
	case rand <= 85:
		gtype = GatesTypeVista
		if len(gl.vistaGatesList) > 0 {
			ng := gl.vistaGatesList[0]
			gl.vistaGatesList = gl.vistaGatesList[1:]
			ng.Reset(x, y)
			return ng, nil
		}
	case rand <= 95:
		gtype = GatesType10
		if len(gl.win10GatesList) > 0 {
			ng := gl.win10GatesList[0]
			gl.win10GatesList = gl.win10GatesList[1:]
			ng.Reset(x, y)
			return ng, nil
		}
	default:
		gtype = GatesTypeXP
		if len(gl.xpGatesList) > 0 {
			ng := gl.xpGatesList[0]
			gl.xpGatesList = gl.xpGatesList[1:]
			ng.Reset(x, y)
			return ng, nil
		}
	}

	ng, err := NewGates(x, y, gtype)
	if err != nil {
		return nil, err
	}

	return ng, nil
}
