package graphics

import "github.com/GomeBox/gome/primitives"

type TextureCreator interface {
	Create(dimensions *primitives.Dimensions, color *primitives.Color) (Texture, error)
}
