package controller

type MouseClickEventHandler interface {
	Do(MouseEvent) error
}
