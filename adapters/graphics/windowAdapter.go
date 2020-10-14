package graphics

//WindowAdapter is to be implemented by the graphics adapter
type WindowAdapter interface {
	//ShowWindow opens a game window
	ShowWindow(windowSettings *WindowSettings) error
}

//WindowSettings contains all the settings needed to show the game window
type WindowSettings struct {
	Fullscreen bool
	Resolution Dimensions
	Title      string
}

//Dimensions defines a width and a height of an entity
type Dimensions struct {
	Width, Height int
}
