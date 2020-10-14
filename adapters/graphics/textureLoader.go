package graphics

type TextureLoader interface {
	Load(fileName string) (TextureDrawer, error)
}
