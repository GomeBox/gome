package mocks

import "github.com/GomeBox/gome/adapters/graphics"

type GraphicsAdapters struct {
}

func (graphics *GraphicsAdapters) TextureLoader() graphics.TextureLoader {
	return nil
}

func (graphics *GraphicsAdapters) FontLoader() graphics.FontLoader {
	return nil
}

func (graphics *GraphicsAdapters) ScreenPresenter() graphics.ScreenPresenter {
	return nil
}

func (graphics *GraphicsAdapters) WindowAdapter() graphics.WindowAdapter {
	return nil
}
