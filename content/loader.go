package content

import "github.com/GomeBox/gome/interfaces"

type Loader interface {
	LoadSound(fileName string) (interfaces.Player, error)
	LoadSong(fileName string) (interfaces.Player, error)
	LoadTexture(fileName string) (interfaces.Texture, error)
	LoadFont(fileName string, size int) (interfaces.Font, error)
}
