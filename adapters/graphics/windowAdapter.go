package graphics

//WindowAdapter is to be implemented by the graphics adapter
type WindowAdapter interface {
	//ShowWindow opens a game window
	ShowWindow(windowSettings *WindowSettings) error
}
