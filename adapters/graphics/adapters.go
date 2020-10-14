package graphics

type Adapters interface {
	TextureLoader() TextureLoader
	WindowAdapter() WindowAdapter
}
