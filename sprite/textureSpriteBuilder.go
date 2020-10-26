package sprite

import (
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/internal/sprite"
	"github.com/GomeBox/gome/primitives"
)

//TextureSpriteBuilder contains functions to create a sprite containing a texture
type TextureSpriteBuilder interface {
	//Get creates the sprite
	Get() Interface
	//SetPosition sets the initial position that the sprite will have after it is created
	SetPosition(x, y float32) TextureSpriteBuilder
	//SetSourceRect can be used to define, which part of the texture should be drawn. If not called or if nil is passed,
	//the whole texture will be drawn
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
	drawer := sprite.NewTextureDrawer(builder.texture, builder.sourceRect)
	s := sprite.New(drawer, builder.pos)
	return s
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
