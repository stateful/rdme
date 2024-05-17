//go:build !windows

package runnerv2service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stateful/runme/v3/internal/project/testdata"
	runnerv2alpha1 "github.com/stateful/runme/v3/pkg/api/gen/proto/go/runme/runner/v2alpha1"
)

// TODO(adamb): add a test case with project.
func TestRunnerServiceSessions(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	t.Run("WithEnv", func(t *testing.T) {
		createResp, err := client.CreateSession(context.Background(), &runnerv2alpha1.CreateSessionRequest{})
		require.NoError(t, err)
		require.NotNil(t, createResp.Session)

		createResp, err = client.CreateSession(context.Background(), &runnerv2alpha1.CreateSessionRequest{Env: []string{"TEST1=value1"}})
		require.NoError(t, err)
		require.EqualValues(t, []string{"TEST1=value1"}, createResp.Session.Env)

		getResp, err := client.GetSession(context.Background(), &runnerv2alpha1.GetSessionRequest{Id: createResp.Session.Id})
		require.NoError(t, err)
		require.EqualValues(t, []string{"TEST1=value1"}, getResp.Session.Env)

		updateResp, err := client.UpdateSession(
			context.Background(),
			&runnerv2alpha1.UpdateSessionRequest{Id: createResp.Session.Id, Env: []string{"TEST2=value2"}},
		)
		require.NoError(t, err)
		require.Equal(t, []string{"TEST1=value1", "TEST2=value2"}, updateResp.Session.Env)

		deleteResp, err := client.DeleteSession(context.Background(), &runnerv2alpha1.DeleteSessionRequest{Id: updateResp.Session.Id})
		require.NoError(t, err)
		require.NotNil(t, deleteResp)

		getResp, err = client.GetSession(context.Background(), &runnerv2alpha1.GetSessionRequest{Id: createResp.Session.Id})
		require.Error(t, err)
		require.Nil(t, getResp)
	})

	t.Run("WithProject", func(t *testing.T) {
		projectPath := testdata.GitProjectPath()
		createResp, err := client.CreateSession(
			context.Background(),
			&runnerv2alpha1.CreateSessionRequest{Project: &runnerv2alpha1.Project{Root: projectPath, EnvLoadOrder: []string{".env"}}},
		)
		require.NoError(t, err)
		require.NotNil(t, createResp.Session)
		require.EqualValues(t, []string{"PROJECT_ENV_FROM_DOTFILE=1"}, createResp.Session.Env)
	})

	t.Run("WithProjectInvalid", func(t *testing.T) {
		_, err := client.CreateSession(
			context.Background(),
			&runnerv2alpha1.CreateSessionRequest{Project: &runnerv2alpha1.Project{Root: "/non/existing/path"}},
		)
		require.Error(t, err)
	})
}

func TestRunnerServiceSessions_StrategyMostRecent(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	// Create a session with env.
	sessResp, err := client.CreateSession(
		context.Background(),
		&runnerv2alpha1.CreateSessionRequest{
			Env: []string{"TEST1=value1"},
		},
	)
	require.NoError(t, err)

	// Prep the execute stream.
	stream, err := client.Execute(context.Background())
	require.NoError(t, err)

	execResult := make(chan executeResult)
	go getExecuteResult(stream, execResult)

	// Execute a program using the most recent session strategy.
	req := &runnerv2alpha1.ExecuteRequest{
		Config: &runnerv2alpha1.ProgramConfig{
			ProgramName: "bash",
			Source: &runnerv2alpha1.ProgramConfig_Commands{
				Commands: &runnerv2alpha1.ProgramConfig_CommandList{
					Items: []string{
						`echo "TEST1=$TEST1"`,
					},
				},
			},
		},
		SessionStrategy: runnerv2alpha1.SessionStrategy_SESSION_STRATEGY_MOST_RECENT,
	}
	err = stream.Send(req)
	require.NoError(t, err)

	result := <-execResult

	assert.NoError(t, result.Err)
	assert.EqualValues(t, 0, result.ExitCode)
	assert.Equal(t, "TEST1=value1\n", string(result.Stdout))

	// Delete the session.
	_, err = client.DeleteSession(context.Background(), &runnerv2alpha1.DeleteSessionRequest{Id: sessResp.GetSession().GetId()})
	require.NoError(t, err)
}
