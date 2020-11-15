package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

type SystemMock struct {
	OnWindow       func() Window
	OnClear        func() error
	CallCntClear   int
	CallCntPresent int
}

func (s *SystemMock) LoadTexture(filename string) (Texture, error) {
	panic("implement me")
}

func (s *SystemMock) LoadFont(fileName string, size int) (Font, error) {
	panic("implement me")
}

func (s *SystemMock) CreateTexture(dimensions primitives.Dimensions, color primitives.Color) (Texture, error) {
	panic("implement me")
}

func (s *SystemMock) Window() Window {
	if s.OnWindow != nil {
		return s.OnWindow()
	}
	return nil
}

func (s *SystemMock) Clear() error {
	s.CallCntClear++
	if s.OnClear != nil {
		return s.OnClear()
	}
	return nil
}

func (s *SystemMock) Present() error {
	s.CallCntPresent++
	return nil
}
