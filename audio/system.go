package audio

import "github.com/GomeBox/gome/internal/audio/interfaces"

//System contains functions to access the audio system
type System interface {
	//LoadSound loads a sound effect from a file
	LoadSound(fileName string) (Player, error)
	//LoadSong loads music from a file
	LoadSong(fileName string) (Player, error)
}

func NewSystem(internalSystem interfaces.System) System {
	return &system{internalSystem: internalSystem}
}

type system struct {
	internalSystem interfaces.System
}

func (sys *system) LoadSound(fileName string) (Player, error) {
	return loadPlayer(fileName, sys.internalSystem.LoadSound)
}

func (sys *system) LoadSong(fileName string) (Player, error) {
	return loadPlayer(fileName, sys.internalSystem.LoadSong)
}

func loadPlayer(fileName string, loadFunc func(fileName string) (interfaces.Player, error)) (*player, error) {
	s, err := loadFunc(fileName)
	if err != nil {
		return nil, err
	}
	p := &player{internalPlayer: s}
	return p, nil
}
