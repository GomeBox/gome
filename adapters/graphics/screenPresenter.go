package graphics

//ScreenPresenter is the adapter that is used to update the game screen
type ScreenPresenter interface {
	//Present updates the screen from the screen buffer
	Present() error
	//Clear deletes the screen buffer
	Clear() error
}
