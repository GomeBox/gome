package graphics

type ScreenPresenter interface {
	Present() error
	Clear() error
}
