package audio

import (
	adapters "github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/interfaces"
)

type System interface {
	LoadSound(fileName string) (interfaces.Player, error)
	LoadSong(fileName string) (interfaces.Player, error)
}

type Adapters struct {
	SoundLoader adapters.SoundLoader
	SongLoader  adapters.SongLoader
}

func NewSystem(adapters Adapters) interfaces.Audio {
	system := system{
		soundLoader: adapters.SoundLoader,
		songLoader:  adapters.SongLoader}
	return &system
}

type system struct {
	soundLoader adapters.SoundLoader
	songLoader  adapters.SongLoader
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
