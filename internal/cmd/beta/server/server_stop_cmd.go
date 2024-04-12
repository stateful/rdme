package server

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/stateful/runme/v3/internal/config"
	"github.com/stateful/runme/v3/internal/config/autoconfig"
)

func serverStopCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "stop",
		Short: "Stop a server.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return autoconfig.Invoke(
				func(
					cfg *config.Config,
					logger *zap.Logger,
				) error {
					defer logger.Sync()

					logger.Debug("stopping the server by looking for runme.pid")

					path := pidFileNameFromAddr(cfg.Kernel.ServerAddress)
					if path == "" {
						return errors.New("server address is not a unix socket")
					}

					pidRaw, err := os.ReadFile(path)
					if err != nil {
						return errors.WithStack(err)
					}
					pid, err := strconv.Atoi(string(pidRaw))
					if err != nil {
						return errors.WithStack(err)
					}

					logger.Debug("found PID file", zap.String("path", path), zap.Int("pid", pid))

					process, err := os.FindProcess(pid)
					if err != nil {
						return errors.WithStack(err)
					}
					return errors.WithStack(process.Signal(os.Interrupt))
				},
			)
		},
	}

	return &cmd
}
