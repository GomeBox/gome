package mocks

import (
	"github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type System struct {
	OnWindow       func() interfaces.Window
	OnClear        func() error
	CallCntClear   int
	CallCntPresent int
}

func (s *System) LoadTexture(filename string) (interfaces.Texture, error) {
	panic("implement me")
}

func (s *System) LoadFont(fileName string, size int) (interfaces.Font, error) {
	panic("implement me")
}

func (s *System) CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (interfaces.Texture, error) {
	panic("implement me")
}

func (s *System) Window() interfaces.Window {
	if s.OnWindow != nil {
		return s.OnWindow()
	}
	return nil
}

func (s *System) Clear() error {
	s.CallCntClear++
	if s.OnClear != nil {
		return s.OnClear()
	}
	return nil
}

func (s *System) Present() error {
	s.CallCntPresent++
	return nil
}
