package interfaces

import "github.com/GomeBox/gome/audio"

type System interface {
	LoadSound(fileName string) (audio.Sound, error)
	LoadSong(fileName string) (audio.Song, error)
}
