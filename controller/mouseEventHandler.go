package controller

// Mouse
type MouseClickEventHandler interface {
	MouseClicked(MouseEvent) error
}
