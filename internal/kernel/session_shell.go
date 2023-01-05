package kernel

import (
	"bytes"
	"io"
	"os"

	"github.com/pkg/errors"
	xpty "github.com/stateful/runme/internal/pty"
	"go.uber.org/zap"
	"golang.org/x/term"
)

type ShellSession struct {
	*Session

	cancelResize  xpty.CancelFn
	stdinOldState *term.State
}

func NewShellSession(
	prompt []byte,
	commandName string,
) (*ShellSession, error) {
	sess, err := NewSession(prompt, commandName, zap.NewNop())
	if err != nil {
		return nil, err
	}

	s := ShellSession{Session: sess}
	s.logger = zap.NewNop()

	s.output = os.Stdout
	s.outputCopyFunc = s.copyOutput

	s.cancelResize = xpty.ResizeOnSig(s.ptmx)

	// Set stdin in raw mode.
	s.stdinOldState, err = term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to put stdin in raw mode")
	}

	go func() {
		// TODO: cancel this goroutine when ptmx is closed
		_, err := io.Copy(s.ptmx, os.Stdin)
		s.setErrorf(err, "failed to copy stdin to ptmx")
	}()

	return &s, nil
}

func (s *ShellSession) Destroy() error {
	s.cancelResize()

	if err := term.Restore(int(os.Stdin.Fd()), s.stdinOldState); err != nil {
		return errors.WithStack(err)
	}

	return s.Session.Destroy()
}

func (s *ShellSession) copyOutput() error {
	buf := make([]byte, 4096)
	if cap(buf) < len(s.prompt) {
		return errors.Errorf("prompt is longer than buffer size (%d vs %d)", len(s.prompt), cap(buf))
	}

	var err error

	for {
		nr, er := s.ptmx.Read(buf)
		if nr > 0 {
			s.outputMu.Lock()
			out := s.output
			s.outputMu.Unlock()

			nw, ew := out.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errors.New("invalid write result")
				}
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = errors.WithStack(io.ErrShortWrite)
				break
			}
		}
		if er != nil {
			err = errors.WithStack(er)
			break
		}
		if bytes.Contains(buf[0:nr], s.prompt) {
			s.promptDetected()
		}
	}
	return err
}
