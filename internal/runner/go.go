package runner

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

type Go struct {
	Base
	Source string
}

var _ Executable = (*Shell)(nil)

func (g Go) DryRun(ctx context.Context, w io.Writer) {
	_, err := exec.LookPath("go")
	if err != nil {
		_, _ = fmt.Fprintf(w, "failed to find %q executable: %s\n", "go", err)
	}

	_, _ = fmt.Fprintf(w, "// go run main.go in $TEMP\n\n")
	_, _ = fmt.Fprintf(w, "%s\n", g.Source)
}

func (g Go) Run(ctx context.Context) error {
	executable, err := exec.LookPath("go")
	if err != nil {
		return errors.Wrapf(err, "failed to find %q executable", "go")
	}

	tmpDir, err := os.MkdirTemp("", "runme-*")
	if err != nil {
		return errors.Wrapf(err, "failed to create a temp dir")
	}
	defer os.RemoveAll(tmpDir)

	mainFile := filepath.Join(tmpDir, "main.go")

	err = os.WriteFile(mainFile, []byte(g.Source), 0o600)
	if err != nil {
		return errors.Wrapf(err, "failed to write source to file")
	}

	c := exec.CommandContext(ctx, executable, "run", mainFile)
	c.Dir = g.Dir
	c.Stderr = g.Stderr
	c.Stdout = g.Stdout
	c.Stdin = g.Stdin

	return errors.Wrapf(c.Run(), "failed to run command %q", "go run main.go")
}
