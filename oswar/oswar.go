package oswar

const (
	IsPlaying = iota
	GameOver
)

const (
	defaultState = IsPlaying
)

const (
	screenSizeWidth  int = 640
	screenSizeHeight int = 480
)

var (
	state = defaultState
)

func SetState(s int) {
	state = s
}

func GetState() int {
	return state
}

func GetScreenWidth() int {
	return screenSizeWidth
}

func GetScreenHeight() int {
	return screenSizeHeight
}