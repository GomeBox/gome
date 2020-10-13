package graphics

type Interface interface {
	ShowWindow(windowSettings *WindowSettings) error
}

type WindowSettings struct {
	Fullscreen bool
	Resolution Dimensions
	Title      string
}

type Dimensions struct {
	Width, Height int
}
