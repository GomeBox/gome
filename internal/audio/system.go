package audio

import (
	"github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/interfaces"
)

func NewSystem(audioAdapters audio.Adapters) interfaces.Audio {
	system := system{
		soundLoader: audioAdapters.SoundLoader(),
		songLoader:  audioAdapters.SongLoader()}
	return &system
}

type system struct {
	soundLoader audio.SoundLoader
	songLoader  audio.SongLoader
}

func (sys *system) LoadSound(fileName string) (interfaces.Player, error) {
	player, err := sys.soundLoader.Load(fileName)
	if err != nil {
		return nil, err
	}
	return &sound{soundPlayer: player}, nil
}

func (sys *system) LoadSong(fileName string) (interfaces.Player, error) {
	player, err := sys.songLoader.Load(fileName)
	if err != nil {
		return nil, err
	}
	return &song{songPlayer: player}, nil
}
