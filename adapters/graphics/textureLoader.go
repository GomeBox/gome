package graphics

//TextureLoader is the adapter that is used to draw a texture
type TextureLoader interface {
	Load(fileName string) (Texture, error)
}
