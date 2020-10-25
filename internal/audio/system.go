package audio

import (
	adapters "github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/audio"
)

type Adapters struct {
	SoundLoader adapters.SoundLoader
	SongLoader  adapters.SongLoader
}

func NewSystem(adapters Adapters) audio.System {
	system := system{
		soundLoader: adapters.SoundLoader,
		songLoader:  adapters.SongLoader}
	return &system
}

type system struct {
	soundLoader adapters.SoundLoader
	songLoader  adapters.SongLoader
}

func (sys *system) LoadSound(fileName string) (audio.Sound, error) {
	player, err := sys.soundLoader.Load(fileName)
	if err != nil {
		return nil, err
	}
	return &sound{soundPlayer: player}, nil
}

func (sys *system) LoadSong(fileName string) (audio.Song, error) {
	player, err := sys.songLoader.Load(fileName)
	if err != nil {
		return nil, err
	}
	return &song{songPlayer: player}, err
}
