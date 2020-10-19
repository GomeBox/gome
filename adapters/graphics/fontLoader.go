package graphics

//FontLoader is the adapter that is used to load fonts
type FontLoader interface {
	//Loads a font from the filesystem and returns a TextCreator
	Load(fileName string, size int) (TextCreator, error)
}
