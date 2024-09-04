//go:build !windows
// +build !windows

package command

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"github.com/stateful/runme/v3/internal/command/testdata"
	runnerv2 "github.com/stateful/runme/v3/pkg/api/gen/proto/go/runme/runner/v2"
)

func TestInlineShellCommand_CollectEnv(t *testing.T) {
	t.Parallel()

	t.Run("Fifo", func(t *testing.T) {
		envCollectorUseFifo = true
		testInlineShellCommandCollectEnv(t)
	})

	t.Run("KillCommandWhileUsingFifo", func(t *testing.T) {
		envCollectorUseFifo = true

		cfg := &ProgramConfig{
			ProgramName: "bash",
			Source: &runnerv2.ProgramConfig_Commands{
				Commands: &runnerv2.ProgramConfig_CommandList{
					Items: []string{
						"export TEST_ENV=1",
						"sleep 5",
					},
				},
			},
			Mode: runnerv2.CommandMode_COMMAND_MODE_INLINE,
		}
		sess := NewSession()
		factory := NewFactory(WithLogger(zaptest.NewLogger(t)))

		command, err := factory.Build(cfg, CommandOptions{Session: sess})
		require.NoError(t, err)
		err = command.Start(context.Background())
		require.NoError(t, err)

		errC := make(chan error, 1)
		go func() {
			<-time.After(time.Second)
			errC <- command.Signal(os.Kill)
		}()
		err = <-errC
		require.NoError(t, err)

		err = command.Wait()
		require.EqualError(t, err, "signal: killed")

		got, ok := sess.GetEnv("TEST_ENV")
		assert.False(t, ok)
		assert.Equal(t, "", got)
	})

	t.Run("NonFifo", func(t *testing.T) {
		envCollectorUseFifo = false
		testInlineShellCommandCollectEnv(t)
	})
}

func TestInlineShellCommand_LimitEnvironSize(t *testing.T) {
	t.Parallel()

	prevMaxSize := MaxEnvironSizeInBytes
	MaxEnvironSizeInBytes = 4096
	defer func() { MaxEnvironSizeInBytes = prevMaxSize }()

	cfg := &ProgramConfig{
		ProgramName: "bash",
		Source: &runnerv2.ProgramConfig_Commands{
			Commands: &runnerv2.ProgramConfig_CommandList{
				Items: []string{
					"echo -n $" + StoreStdoutEnvName,
				},
			},
		},
		Mode: runnerv2.CommandMode_COMMAND_MODE_FILE,
	}

	factory := NewFactory(
		WithLogger(zaptest.NewLogger(t)),
		WithRuntime(&hostRuntime{}), // stub runtime and do not include environ
	)

	sess := NewSession()
	err := sess.SetEnv(
		StoreStdoutEnvName + "=" + strings.Repeat("a", MaxEnvironSizeInBytes),
	)
	require.NoError(t, err)

	stdout := bytes.NewBuffer(nil)

	command, err := factory.Build(cfg, CommandOptions{Session: sess, Stdout: stdout})
	require.NoError(t, err)
	err = command.Start(context.Background())
	require.NoError(t, err)
	err = command.Wait()
	require.NoError(t, err)
	expectedLen := MaxEnvironSizeInBytes - len(StoreStdoutEnvName) - 1 // subtract len("<KEY>=")
	require.Equal(t, expectedLen, stdout.Len())
}

func TestInlineShellCommand_LargeOutput(t *testing.T) {
	t.Parallel()

	temp := t.TempDir()
	fileName := filepath.Join(temp, "large_output.json.gzip")
	err := os.WriteFile(fileName, testdata.Users1MGzip, 0o644)
	require.NoError(t, err)

	factory := NewFactory(WithLogger(zaptest.NewLogger(t)))
	sess := NewSession()

	cfg := &ProgramConfig{
		ProgramName: "bash",
		Source: &runnerv2.ProgramConfig_Commands{
			Commands: &runnerv2.ProgramConfig_CommandList{
				Items: []string{
					"cat " + fileName,
				},
			},
		},
		Mode: runnerv2.CommandMode_COMMAND_MODE_INLINE,
	}

	stdout := bytes.NewBuffer(nil)

	command, err := factory.Build(cfg, CommandOptions{Session: sess, Stdout: stdout})
	require.NoError(t, err)

	err = command.Start(context.Background())
	require.NoError(t, err)
	err = command.Wait()
	require.NoError(t, err)

	require.NoError(t, err)
	assert.EqualValues(t, testdata.Users1MGzip, stdout.Bytes())
}

func testInlineShellCommandCollectEnv(t *testing.T) {
	t.Helper()

	cfg := &ProgramConfig{
		ProgramName: "bash",
		Source: &runnerv2.ProgramConfig_Script{
			Script: "export TEST_ENV=1",
		},
		Mode: runnerv2.CommandMode_COMMAND_MODE_INLINE,
	}
	sess := NewSession()

	testExecuteCommandWithSession(t, cfg, sess, nil, "", "")

	got, ok := sess.GetEnv("TEST_ENV")
	assert.True(t, ok)
	assert.Equal(t, "1", got)
}
