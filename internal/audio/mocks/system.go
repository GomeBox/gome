package mocks

import (
	"github.com/GomeBox/gome/internal/audio/interfaces"
)

type System struct {
}

func (s System) LoadSound(fileName string) (interfaces.Sound, error) {
	panic("implement me")
}

func (s System) LoadSong(fileName string) (interfaces.Song, error) {
	panic("implement me")
}
