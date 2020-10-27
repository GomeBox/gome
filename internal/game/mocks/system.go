package mocks

import (
	"github.com/GomeBox/gome/internal/audio"
	"github.com/GomeBox/gome/internal/graphics"
	"github.com/GomeBox/gome/internal/input"
)

type System struct {
	OnGraphics func() *graphics.System
	OnInput    func() *input.System
	OnAudio    func() *audio.System
}

func (s System) Initialize() error {
	panic("implement me")
}

func (s System) Update() error {
	panic("implement me")
}

func (s System) Graphics() *graphics.System {
	if s.OnGraphics != nil {
		return s.OnGraphics()
	}
	return nil
}

func (s System) Input() *input.System {
	if s.OnInput != nil {
		return s.OnInput()
	}
	return nil
}

func (s System) Audio() *audio.System {
	if s.OnAudio != nil {
		return s.OnAudio()
	}
	return nil
}
