package graphics

import "github.com/GomeBox/gome/adapters/graphics"

type Adapters struct {
	TextureLoader graphics.TextureLoader
	WindowAdapter graphics.WindowAdapter
}

type System interface {
	LoadTexture(fileName string) (Texture, error)
}

func NewSystem(adapters Adapters) System {
	sys := new(system)
	sys.textureLoader = adapters.TextureLoader
	return sys
}

type system struct {
	textureLoader graphics.TextureLoader
	windowAdapter graphics.WindowAdapter
}

func (sys *system) LoadTexture(filename string) (Texture, error) {
	drawer, err := sys.textureLoader.Load(filename)
	if err != nil {
		return nil, err
	}
	return NewTexture(drawer), nil
}
