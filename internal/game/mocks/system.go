package mocks

import (
	audio "github.com/GomeBox/gome/internal/audio/interfaces"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
	input "github.com/GomeBox/gome/internal/input/interfaces"
)

type System struct {
	OnGraphics func() graphics.System
	OnInput    func() input.System
	OnAudio    func() audio.System
}

func (s System) Initialize() error {
	panic("implement me")
}

func (s System) Update() error {
	panic("implement me")
}

func (s System) Graphics() graphics.System {
	if s.OnGraphics != nil {
		return s.OnGraphics()
	}
	return nil
}

func (s System) Input() input.System {
	if s.OnInput != nil {
		return s.OnInput()
	}
	return nil
}

func (s System) Audio() audio.System {
	if s.OnAudio != nil {
		return s.OnAudio()
	}
	return nil
}
