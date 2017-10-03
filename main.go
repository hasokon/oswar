package main

import (
	"github.com/hasokon/oswar/oswar"
	"github.com/hasokon/oswar/oswar/view"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.Run(view.Update(), oswar.GetScreenWidth(), oswar.GetScreenHeight(), 1, "OS War")
}
