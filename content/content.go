package content

import "github.com/GomeBox/gome/interfaces"

type Content interface {
	Loader
	interfaces.Unloader
}

func New(context interfaces.Context) Content {
	c := new(content)
	c.context = context
	c.loadedContent = make([]interfaces.Unloader, 0)
	return c
}

type content struct {
	context       interfaces.Context
	loadedContent []interfaces.Unloader
}

func (c *content) LoadSound(fileName string) (interfaces.Player, error) {
	player, err := c.context.Audio().LoadSound(fileName)
	c.loadedContent = append(c.loadedContent, player)
	return player, err
}

func (c *content) LoadSong(fileName string) (interfaces.Player, error) {
	player, err := c.context.Audio().LoadSong(fileName)
	c.loadedContent = append(c.loadedContent, player)
	return player, err
}

func (c *content) LoadTexture(fileName string) (interfaces.Texture, error) {
	panic("implement me")
}

func (c *content) LoadFont(fileName string, size int) (interfaces.Font, error) {
	panic("implement me")
}

func (c *content) Unload() error {
	panic("implement me")
}

func (c *content) Unloaded() bool {
	return false
}
