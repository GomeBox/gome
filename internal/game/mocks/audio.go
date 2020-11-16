package mocks

import (
	"github.com/GomeBox/gome/interfaces"
)

type Audio struct {
}

func (a Audio) LoadSound(fileName string) (interfaces.Player, error) {
	panic("implement me")
}

func (a Audio) LoadSong(fileName string) (interfaces.Player, error) {
	panic("implement me")
}
