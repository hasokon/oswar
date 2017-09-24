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
	return &MouseEventListener{
		lastMousePressed:        false,
		currentMousePressed:     false,
		mouseClickEventHandlers: make([]MouseClickEventHandler, 0),
	}
}

func (ml *MouseEventListener) isMouseClicked() bool {
	ml.lastMousePressed = ml.currentMousePressed
	ml.currentMousePressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	return !ml.lastMousePressed && ml.currentMousePressed
}

func (ml *MouseEventListener) AddMouseClickEventHandler(m MouseClickEventHandler) {
	ml.mouseClickEventHandlers = append(ml.mouseClickEventHandlers, m)
}

func (ml *MouseEventListener) Update() {
	x, y := ebiten.CursorPosition()
	me := MouseEvent{x, y}
	switch {
	case ml.isMouseClicked():
		for _, handler := range ml.mouseClickEventHandlers {
			handler.MouseClicked(me)
		}
	}
}
