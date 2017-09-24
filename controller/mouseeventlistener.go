package controller

import (
	"github.com/hajimehoshi/ebiten"
)

type MouseEventListener struct {
	lastMousePressed        bool
	currentMousePressed     bool
	mouseClickEventHandlers []MouseClickEventHandler
}

func New() *MouseEventListener {
	oc := MouseEventListener{
		lastMousePressed:        false,
		currentMousePressed:     false,
		mouseClickEventHandlers: make([]MouseClickEventHandler, 0),
	}
	return &oc
}

func (ml *MouseEventListener) isMouseClicked() bool {
	ml.lastMousePressed = ml.currentMousePressed
	ml.currentMousePressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	return !ml.lastMousePressed && ml.currentMousePressed
}

func (ml *MouseEventListener) AddMouseClickEventHandler(m MouseClickEventHandler) {
	ml.mouseClickEventHandlers = append(ml.mouseClickEventHandlers, m)
}

func (oc *MouseEventListener) Update() {
	x, y := ebiten.CursorPosition()
	me := MouseEvent{x, y}
	switch {
	case oc.isMouseClicked():
		for i := 0; i < len(oc.mouseClickEventHandlers); i++ {
			oc.mouseClickEventHandlers[i].Do(me)
		}
	}
}
