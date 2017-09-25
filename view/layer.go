package view

import "github.com/hajimehoshi/ebiten"

type Layer interface {
	Update() error
	Canvas() *ebiten.Image
}
