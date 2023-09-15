package executor_test

import (
	"runtime"
	"strings"
	"testing"

	"github.com/jithinkunjachan/nasserver/backend/pkg/executor"
	"github.com/stretchr/testify/assert"
)

type mockCommand struct {
}

var _ executor.Builder = &mockCommand{}
var isError = false

func (me *mockCommand) Command() executor.Command {

	if isError {
		return executor.Command{
			Cmd:  "wrongcomand",
			Args: nil,
			Ws:   nil,
		}
	}
	cmd := executor.Command{
		Cmd:  "ls",
		Args: nil,
		Ws:   nil,
	}
	if runtime.GOOS == "windows" {
		cmd = executor.Command{
			Cmd:  "cmd",
			Args: []string{"/C", "dir"},
			Ws:   nil,
		}
	}
	return cmd
}

func Test_executor(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		// given
		isError = false
		defer func() {
			isError = false
		}()

		// when
		result, err := executor.Exec(&mockCommand{})

		//then
		assert.True(t, strings.ContainsAny(result, "go.sum"))
		assert.True(t, strings.ContainsAny(result, "go.mod"))
		assert.NoError(t, err)
	})
	t.Run("error case should have empty result", func(t *testing.T) {
		// given
		isError = true
		defer func() {
			isError = false
		}()

		// when
		result, err := executor.Exec(&mockCommand{})

		//then
		assert.Equal(t, "", result)
		assert.Error(t, err)
	})
}
