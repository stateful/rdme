package command

import (
	"context"
	"io"
	"os"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type terminalCommand struct {
	internalCommand

	envCollector envCollector
	logger       *zap.Logger
	session      *Session
	stdinWriter  io.Writer
}

func (c *terminalCommand) getPty() *os.File {
	cmdPty, ok := c.internalCommand.(commandWithPty)
	if !ok {
		return nil
	}
	return cmdPty.getPty()
}

func (c *terminalCommand) Start(ctx context.Context) (err error) {
	if isNil(c.stdinWriter) {
		return errors.New("stdin writer is nil")
	}

	c.logger.Info("starting a terminal command")
	if err := c.internalCommand.Start(ctx); err != nil {
		return err
	}
	c.logger.Info("a terminal command started")

	// [shellEnvCollector] writes defines a function collecting env and
	// registers it as a trap directly into the shell interactive session.
	opts := envCollectorFactoryOptions{
		useFifo: useEnvCollectorFifo,
	}
	c.envCollector, err = newEnvCollectorFactory(opts).Build()
	if err != nil {
		return err
	}
	return c.envCollector.SetOnShell(c.stdinWriter)
}

func (c *terminalCommand) Wait() (err error) {
	err = c.internalCommand.Wait()
	if cErr := c.collectEnv(); cErr != nil {
		c.logger.Info("failed to collect the environment", zap.Error(cErr))
		if err == nil {
			err = cErr
		}
	}
	return err
}

func (c *terminalCommand) collectEnv() error {
	if c.session == nil || c.envCollector == nil {
		return nil
	}

	changed, deleted, err := c.envCollector.Diff()
	if err != nil {
		return err
	}

	err = c.session.SetEnv(changed...)
	if err != nil {
		return errors.WithMessage(err, "failed to set the new or updated env")
	}

	c.session.DeleteEnv(deleted...)

	return nil
}
