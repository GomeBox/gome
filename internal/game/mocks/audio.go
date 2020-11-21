package mocks

import (
	"github.com/GomeBox/gome/interfaces"
)

type Audio struct {
	OnLoadSound func(fileName string) (interfaces.Player, error)
	OnLoadSong  func(fileName string) (interfaces.Player, error)
}

func (a *Audio) LoadSound(fileName string) (interfaces.Player, error) {
	if a.OnLoadSound != nil {
		return a.OnLoadSound(fileName)
	}
	return nil, nil
}

func (a *Audio) LoadSong(fileName string) (interfaces.Player, error) {
	if a.OnLoadSong != nil {
		return a.OnLoadSong(fileName)
	}
	return nil, nil
}
