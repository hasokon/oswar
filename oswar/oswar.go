package oswar

import (
	"github.com/hasokon/oswar/oswar/view"
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	isPlaying = iota
)

const (
	defaultState = isPlaying
)

const (
	screenSizeWidth  int = 640
	screenSizeHeight int = 480
)

var (
	state = defaultState
)

func GetState() int {
	return state
}

func GetScreenWidth() int {
	return screenSizeWidth
}

func GetScreenHeight() int {
	return screenSizeHeight
}

func Update() func(*ebiten.Image) error {
	o, _ := view.NewOswar(screenSizeWidth, screenSizeHeight)

	return func(screen *ebiten.Image) error {
		if ebiten.IsRunningSlowly() {
			return nil
		}

		o.Update(screen)

		x, y := ebiten.CursorPosition()
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS : %0.2f\n(%d, %d)", ebiten.CurrentFPS(), x, y))
		return nil
	}
}