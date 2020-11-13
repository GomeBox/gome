package mocks

import (
	"github.com/GomeBox/gome/internal/audio/interfaces"
)

type System struct {
	CallCntLoadSound int
	CallCntLoadSong  int
}

func (s *System) LoadSound(fileName string) (interfaces.Player, error) {
	s.CallCntLoadSound++
	return nil, nil
}

func (s *System) LoadSong(fileName string) (interfaces.Player, error) {
	s.CallCntLoadSong++
	return nil, nil
}
