package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hasokon/oswar/view"
)

const (
	screenSizeWidth  int = 640
	screenSizeHeight int = 480
)

func main() {
	oswar := view.NewOswar(screenSizeWidth, screenSizeHeight)
	ebiten.Run(oswar.GetUpdate(), screenSizeWidth, screenSizeHeight, 1, "Test")
}
