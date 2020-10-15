package graphics

type FontLoader interface {
	Load(fileName string, size int) (TextCreator, error)
}
