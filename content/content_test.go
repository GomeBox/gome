package content

import (
	"errors"
	"github.com/GomeBox/gome/interfaces"
	interfacesMocks "github.com/GomeBox/gome/interfaces/mocks"
	"github.com/GomeBox/gome/internal/game"
	"github.com/GomeBox/gome/internal/game/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	context := game.NewContext(nil, nil, nil)
	testContent := New(context)
	c := testContent.(*content)
	assert.Same(t, c.context, context)
}

func TestContent_LoadSound(t *testing.T) {
	testLoadPlayer(t, Content.LoadSound)
}

func TestContent_LoadSong(t *testing.T) {
	testLoadPlayer(t, Content.LoadSong)
}

func TestContent_LoadTexture(t *testing.T) {
	var gotFilename, wantFilename string
	wantFilename = "test.png"
	retErr := false
	texture := &interfacesMocks.Texture{}
	graphics := mocks.Graphics{
		OnLoadTexture: func(fileName string) (interfaces.Texture, error) {
			gotFilename = fileName
			if retErr {
				return nil, errors.New("test")
			}
			return texture, nil
		},
	}
	context := game.NewContext(&graphics, nil, nil)
	testContent := New(context)
	gotTexture, err := testContent.LoadTexture(wantFilename)
	assert.NoError(t, err)
	assert.Same(t, texture, gotTexture)
	assert.Equal(t, wantFilename, gotFilename)

	testContent = New(context)
	loadFunc := func() {
		_, _ = testContent.LoadTexture(wantFilename)
	}
	testAddedToLoadedContent(t, loadFunc, testContent)
}

func TestContent_LoadFont(t *testing.T) {
	var gotFilename, wantFilename string
	var gotSize, wantSize int
	wantFilename = "test.font"
	wantSize = 12
	retErr := false
	font := &interfacesMocks.Font{}
	graphics := mocks.Graphics{
		OnLoadFont: func(fileName string, size int) (interfaces.Font, error) {
			gotFilename = fileName
			gotSize = size
			if retErr {
				return nil, errors.New("test")
			}
			return font, nil
		},
	}
	context := game.NewContext(&graphics, nil, nil)
	testContent := New(context)
	gotTexture, err := testContent.LoadFont(wantFilename, wantSize)
	assert.NoError(t, err)
	assert.Same(t, font, gotTexture)
	assert.Equal(t, wantFilename, gotFilename)
	assert.Equal(t, wantSize, gotSize)

	testContent = New(context)
	loadFunc := func() {
		_, _ = testContent.LoadTexture(wantFilename)
	}
	testAddedToLoadedContent(t, loadFunc, testContent)
}

func TestContent_Unload(t *testing.T) {
	c := new(content)
	size := 10
	isUnloaded := make([]bool, size)
	for i := 0; i < size; i++ {
		isUnloaded[i] = false
		index := i
		unloader := &interfacesMocks.Unloader{
			OnUnload: func() error {
				isUnloaded[index] = true
				return nil
			},
		}
		c.addLoadedContent(unloader)
	}
	// Test all elements of c.loaded content get deleted
	err := c.Unload()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(c.loadedContent), "not all elements deleted")
	for i := 0; i < size; i++ {
		if !assert.True(t, isUnloaded[i], i) {
			break
		}
	}

	//Test that another call to Unload() does not return an error or panic
	err = c.Unload()
	assert.NoError(t, err)

	// Test that content that was manually unloaded gets deleted from loaded content
	unloadCalled := false
	content := interfacesMocks.Unloader{
		OnUnload: func() error {
			unloadCalled = true
			return nil
		},
		OnUnloaded: func() bool {
			return true
		},
	}
	c.addLoadedContent(&content)
	err = c.Unload()
	assert.NoError(t, err)
	assert.Nil(t, c.loadedContent)
	assert.False(t, unloadCalled)

	// Test that an error is returned if unloading fails
	content = interfacesMocks.Unloader{
		OnUnload: func() error {
			return errors.New("test")
		},
	}
	c.addLoadedContent(&content)
	err = c.Unload()
	assert.Error(t, err)
}

func TestContent_Unloaded(t *testing.T) {
	c := new(content)
	assert.True(t, c.Unloaded())

	cont1 := interfacesMocks.Unloader{
		OnUnloaded: func() bool {
			return true
		},
	}
	c.addLoadedContent(&cont1)
	assert.True(t, c.Unloaded())

	cont2 := interfacesMocks.Unloader{
		OnUnloaded: func() bool {
			return false
		},
	}
	c.addLoadedContent(&cont2)
	assert.False(t, c.Unloaded())
}

func testLoadPlayer(t *testing.T, load func(content Content, filename string) (interfaces.Player, error)) {
	var gotFilename, wantFilename string
	wantFilename = "test.mp3"
	retErr := false
	audio := mocks.Audio{}
	player := new(interfacesMocks.Player)
	audio.OnLoadSound = func(fileName string) (interfaces.Player, error) {
		gotFilename = fileName
		if retErr {
			return nil, errors.New("test")
		}
		return player, nil
	}
	audio.OnLoadSong = func(fileName string) (interfaces.Player, error) {
		gotFilename = fileName
		if retErr {
			return nil, errors.New("test")
		}
		return player, nil
	}
	context := game.NewContext(nil, nil, &audio)
	testContent := New(context)
	gotPlayer, err := load(testContent, wantFilename)
	assert.NoError(t, err)
	assert.Same(t, player, gotPlayer)
	assert.Equal(t, wantFilename, gotFilename)

	testContent = New(context)
	loadFunc := func() {
		_, _ = testContent.LoadSound(wantFilename)
	}
	testAddedToLoadedContent(t, loadFunc, testContent)
}

func testAddedToLoadedContent(t *testing.T, loadFunc func(), testContent Content) {
	c := testContent.(*content)
	for i := 1; i <= 10; i++ {
		loadFunc()
		assert.Equal(t, i, len(c.loadedContent))
	}
}
