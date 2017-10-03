package main

import (
	"github.com/hasokon/oswar/oswar"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.Run(oswar.Update(), oswar.GetScreenWidth(), oswar.GetScreenHeight(), 1, "OS War")
}
