package command

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type pathNormalizer struct {
	kernel Kernel
}

func newPathNormalizer(kernel Kernel) configNormalizer {
	return (&pathNormalizer{kernel: kernel}).Normalize
}

func (n *pathNormalizer) Normalize(cfg *Config) (func() error, error) {
	var (
		programPath string   = cfg.ProgramName
		args        []string = cfg.Arguments
		err         error
	)

	programPath, err = n.kernel.LookPath(cfg.ProgramName)
	if programPath != "" && err == nil {
		goto finish
	}

	if isShellLanguage(cfg.LanguageId) {
		globalShell := shellFromShellPath(globalShellPath())
		programPath, err = n.kernel.LookPath(globalShell)
		if err != nil {
			goto finish
		}
	}

	programPath, args, err = n.findProgramInInterpreters(cfg.ProgramName, cfg.LanguageId)
	if err == nil {
		goto finish
	}

	// Default to "cat" for shebang++
	programPath, err = n.kernel.LookPath("cat")
	args = []string{}
	if err == nil {
		goto finish
	}

	return nil, err

finish:
	cfg.ProgramName = programPath
	cfg.Arguments = args

	return nil, nil
}

func (n *pathNormalizer) findProgramInInterpreters(programName string, languageID string) (programPath string, args []string, _ error) {
	if programName != "" {
		res, err := n.kernel.LookPath(programName)
		if err != nil {
			return "", nil, errors.Errorf("unsupported interpreter %s", programName)
		}
		return res, []string{}, nil
	}

	interpreters := inferInterpreterFromLanguage(languageID)
	if len(interpreters) == 0 {
		return "", nil, errors.Errorf("unsupported interpreter for %s", languageID)
	}

	for _, interpreter := range interpreters {
		iProgram, iArgs := parseInterpreter(interpreter)
		if path, err := n.kernel.LookPath(iProgram); err == nil {
			programPath = path
			args = iArgs
			return
		}
	}

	return "", nil, errors.Errorf("unable to look up any of interpreters %s", interpreters)
}

// parseInterpreter handles cases when the interpreter is, for instance, "deno run".
// Only the first word is a program name and the rest is arguments.
func parseInterpreter(interpreter string) (program string, args []string) {
	parts := strings.SplitN(interpreter, " ", 2)

	if len(parts) > 0 {
		program = parts[0]
	}

	if len(parts) > 1 {
		args = strings.Split(parts[1], " ")
	}

	return
}

func inferInterpreterFromLanguage(langID string) []string {
	return interpreterByLanguageID[langID]
}

var interpreterByLanguageID = map[string][]string{
	"js":              {"node"},
	"javascript":      {"node"},
	"jsx":             {"node"},
	"javascriptreact": {"node"},

	"ts":              {"ts-node", "deno run", "bun run"},
	"typescript":      {"ts-node", "deno run", "bun run"},
	"tsx":             {"ts-node", "deno run", "bun run"},
	"typescriptreact": {"ts-node", "deno run", "bun run"},

	"sh":          {"bash", "sh"},
	"bash":        {"bash", "sh"},
	"ksh":         {"ksh"},
	"zsh":         {"zsh"},
	"fish":        {"fish"},
	"powershell":  {"powershell"},
	"cmd":         {"cmd"},
	"dos":         {"cmd"},
	"shellscript": {"bash", "sh"},

	"lua":    {"lua"},
	"perl":   {"perl"},
	"php":    {"php"},
	"python": {"python3", "python"},
	"py":     {"python3", "python"},
	"ruby":   {"ruby"},
	"rb":     {"ruby"},
}

func globalShellPath() string {
	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		shell = "sh"
	}
	if path, err := exec.LookPath(shell); err == nil {
		return path
	}
	return "/bin/sh"
}

// TODO(sebastian): this method for determining shell is not strong, since shells can
// be aliased. we should probably run the shell to get this information
func shellFromShellPath(programPath string) string {
	programFile := filepath.Base(programPath)
	return programFile[:len(programFile)-len(filepath.Ext(programFile))]
}
