package sprite

import (
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

type TextureSpriteBuilder interface {
	Get() Interface
	SetPosition(x, y float32) TextureSpriteBuilder
	SetSourceRect(sourceRect *primitives.Rectangle) TextureSpriteBuilder
}

func FromTexture(texture graphics.Texture) TextureSpriteBuilder {
	builder := new(textureSpriteBuilder)
	builder.texture = texture
	return builder
}

type textureSpriteBuilder struct {
	texture    graphics.Texture
	pos        primitives.PointF
	sourceRect *primitives.Rectangle
}

func (builder *textureSpriteBuilder) Get() Interface {
	drawer := newTextureDrawer(builder.texture)
	drawer.sourceRect = builder.sourceRect
	sprite := sprite{
		drawer:   drawer,
		position: builder.pos,
	}
	return &sprite
}

func (builder *textureSpriteBuilder) SetPosition(x, y float32) TextureSpriteBuilder {
	builder.pos.X = x
	builder.pos.Y = y
	return builder
}

func (builder *textureSpriteBuilder) SetSourceRect(sourceRect *primitives.Rectangle) TextureSpriteBuilder {
	builder.sourceRect = sourceRect
	return builder
}
