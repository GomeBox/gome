package audio

import "github.com/GomeBox/gome/internal/audio/interfaces"

//System contains functions to access the audio system
type System interface {
	//LoadSound loads a sound effect from a file
	LoadSound(fileName string) (Sound, error)
	//LoadSong loads music from a file
	LoadSong(fileName string) (Song, error)
}

func NewSystem(internalSystem interfaces.System) System {
	return &system{internalSystem: internalSystem}
}

type system struct {
	internalSystem interfaces.System
}

func (sys *system) LoadSound(fileName string) (Sound, error) {
	s, err := sys.internalSystem.LoadSound(fileName)
	if err != nil {
		return nil, err
	}
	return &sound{internalSound: s}, nil
}

func (sys *system) LoadSong(fileName string) (Song, error) {
	s, err := sys.internalSystem.LoadSong(fileName)
	if err != nil {
		return nil, err
	}
	return &song{internalSong: s}, nil
}
