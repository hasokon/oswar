package controller

import (
	"github.com/hajimehoshi/ebiten"
)

type OswarController struct {
	lastMousePressed        bool
	currentMousePressed     bool
	mouseClickEventHandlers []MouseClickEventHandler
}

func New() *OswarController {
	oc := OswarController{
		lastMousePressed:        false,
		currentMousePressed:     false,
		mouseClickEventHandlers: make([]MouseClickEventHandler, 0),
	}
	return &oc
}

func (oc *OswarController) isMouseClicked() bool {
	oc.lastMousePressed = oc.currentMousePressed
	oc.currentMousePressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	return !oc.lastMousePressed && oc.currentMousePressed
}

func (oc *OswarController) AddMouseClickEventHandler(m MouseClickEventHandler) {
	oc.mouseClickEventHandlers = append(oc.mouseClickEventHandlers, m)
}

func (oc *OswarController) Update() {
	x, y := ebiten.CursorPosition()
	me := MouseEvent{x, y}
	switch {
	case oc.isMouseClicked():
		for i := 0; i < len(oc.mouseClickEventHandlers); i++ {
			oc.mouseClickEventHandlers[i].Do(me)
		}
	}
}
