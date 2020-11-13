package mocks

import "github.com/GomeBox/gome/adapters/graphics"

type Adapters struct {
	OnTextureLoader        func() graphics.TextureLoader
	OnTextureCreator       func() graphics.TextureCreator
	OnFontLoader           func() graphics.FontLoader
	OnWindowAdapter        func() graphics.WindowAdapter
	OnScreenPresenter      func() graphics.ScreenPresenter
	CallCntTextureLoader   int
	CallCntTextureCreator  int
	CallCntFontLoader      int
	CallCntWindowAdapter   int
	CallCntScreenPresenter int
}

func (a *Adapters) TextureLoader() graphics.TextureLoader {
	a.CallCntTextureLoader++
	if a.OnTextureLoader != nil {
		return a.OnTextureLoader()
	}
	return nil
}

func (a *Adapters) TextureCreator() graphics.TextureCreator {
	a.CallCntTextureCreator++
	if a.OnTextureCreator != nil {
		return a.OnTextureCreator()
	}
	return nil
}

func (a *Adapters) FontLoader() graphics.FontLoader {
	a.CallCntFontLoader++
	if a.OnFontLoader != nil {
		return a.OnFontLoader()
	}
	return nil
}

func (a *Adapters) WindowAdapter() graphics.WindowAdapter {
	a.CallCntWindowAdapter++
	if a.OnWindowAdapter != nil {
		return a.OnWindowAdapter()
	}
	return nil
}

func (a *Adapters) ScreenPresenter() graphics.ScreenPresenter {
	a.CallCntScreenPresenter++
	if a.OnScreenPresenter != nil {
		return a.OnScreenPresenter()
	}
	return nil
}
