package graphics

type Adapters interface {
	TextureLoader() TextureLoader
	FontLoader() FontLoader
	WindowAdapter() WindowAdapter
	ScreenPresenter() ScreenPresenter
}
