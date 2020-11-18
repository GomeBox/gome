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
	assert.NotNil(t, c.loadedContent)
}

func TestContent_LoadSound(t *testing.T) {
	testLoadPlayer(t, Content.LoadSound)
}

func TestContent_LoadSong(t *testing.T) {
	testLoadPlayer(t, Content.LoadSong)
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
