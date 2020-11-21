package content

import "github.com/GomeBox/gome/interfaces"

type Content interface {
	Loader
	interfaces.Unloader
}

func New(context interfaces.Context) Content {
	c := new(content)
	c.context = context
	return c
}

type content struct {
	context       interfaces.Context
	loadedContent []interfaces.Unloader
}

func (c *content) LoadSound(fileName string) (interfaces.Player, error) {
	player, err := c.context.Audio().LoadSound(fileName)
	c.addLoadedContent(player)
	return player, err
}

func (c *content) LoadSong(fileName string) (interfaces.Player, error) {
	player, err := c.context.Audio().LoadSong(fileName)
	c.addLoadedContent(player)
	return player, err
}

func (c *content) LoadTexture(fileName string) (interfaces.Texture, error) {
	texture, err := c.context.Graphics().LoadTexture(fileName)
	c.addLoadedContent(texture)
	return texture, err
}

func (c *content) LoadFont(fileName string, size int) (interfaces.Font, error) {
	font, err := c.context.Graphics().LoadFont(fileName, size)
	c.addLoadedContent(font)
	return font, err
}

func (c *content) Unload() error {
	var err error
	for _, cont := range c.loadedContent {
		if cont.Unloaded() {
			continue
		}
		tmpErr := cont.Unload()
		if tmpErr != nil {
			err = tmpErr
		}
	}
	if err == nil {
		c.loadedContent = nil
	}
	return err
}

func (c *content) Unloaded() bool {
	allUnloaded := true
	for _, cont := range c.loadedContent {
		if !cont.Unloaded() {
			allUnloaded = false
			break
		}
	}
	return allUnloaded
}

func (c *content) addLoadedContent(content interfaces.Unloader) {
	if c.loadedContent == nil {
		c.loadedContent = make([]interfaces.Unloader, 0)
	}
	c.loadedContent = append(c.loadedContent, content)
}
