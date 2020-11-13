package audio

import (
	"errors"
	"github.com/GomeBox/gome/internal/audio/interfaces"
	"github.com/GomeBox/gome/internal/audio/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSystem(t *testing.T) {
	intSys := &mocks.System{}
	sys := NewSystem(intSys)
	assert.NotEmpty(t, sys)
	s := sys.(*system)
	assert.Same(t, intSys, s.internalSystem)
}

func Test_loadPlayer(t *testing.T) {
	fileName := "test.file"
	var gotFileName string
	retErr := false
	internalPlayer := new(mocks.Player)
	loadFunc := func(fileName string) (interfaces.Player, error) {
		if retErr {
			return nil, errors.New("test")
		}
		gotFileName = fileName
		return internalPlayer, nil
	}
	player, err := loadPlayer(fileName, loadFunc)
	assert.NoError(t, err)
	assert.NotEmpty(t, player)
	assert.Same(t, internalPlayer, player.internalPlayer)
	assert.Equal(t, fileName, gotFileName)
	retErr = true
	_, err = loadPlayer(fileName, loadFunc)
	assert.Error(t, err)
}

func TestSystem_LoadSong(t *testing.T) {
	intSys := new(mocks.System)
	sys := system{internalSystem: intSys}
	player, err := sys.LoadSong("test")
	assert.NoError(t, err)
	assert.NotNil(t, player, "player")
	assert.Equal(t, 1, intSys.CallCntLoadSong, "LoadSong was not called expected number of times")
}

func TestSystem_LoadSound(t *testing.T) {
	intSys := new(mocks.System)
	sys := system{internalSystem: intSys}
	player, err := sys.LoadSound("test")
	assert.NoError(t, err)
	assert.NotNil(t, player, "player")
	assert.Equal(t, 1, intSys.CallCntLoadSound, "LoadSound was not called expected number of times")
}
