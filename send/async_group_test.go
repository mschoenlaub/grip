package send

import (
	"context"
	"testing"

	"github.com/mschoenlaub/grip/level"
	"github.com/mschoenlaub/grip/message"
	"github.com/stretchr/testify/assert"
)

func TestAsyncGroupSender(t *testing.T) {
	assert := assert.New(t) // nolint
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cs := MakeErrorLogger()
	assert.NoError(cs.SetLevel(LevelInfo{Default: level.Debug, Threshold: level.Notice}))
	assert.True(cs.Level().Valid())

	s := NewAsyncGroupSender(ctx, 2, cs)

	// if it's not valid to start with then we shouldn't make it
	// valid by setting it to avoid constituents being overridden,
	assert.False(s.Level().Valid())
	assert.NoError(s.SetLevel(LevelInfo{Default: level.Info, Threshold: level.Alert}))
	assert.False(s.Level().Valid())
	assert.True(cs.Level().Valid())

	impl, ok := s.(*asyncGroupSender)

	assert.True(ok)
	newLevel := LevelInfo{Default: level.Info, Threshold: level.Alert}
	assert.NotEqual(newLevel, s.Level())
	impl.level = newLevel
	assert.Equal(newLevel, s.Level())

	s.Send(message.NewDefaultMessage(level.Debug, "hello"))
	newLevel = LevelInfo{Default: level.Debug, Threshold: level.Alert}
	assert.NoError(impl.SetLevel(newLevel))
	assert.Equal(newLevel, s.Level())

	assert.NoError(s.Close())
}
