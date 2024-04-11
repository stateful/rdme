package beta

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/cli/go-gh/v2/pkg/jsonpretty"
	"github.com/cli/go-gh/v2/pkg/tableprinter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/stateful/runme/v3/internal/command"
	"github.com/stateful/runme/v3/internal/config/autoconfig"
	"github.com/stateful/runme/v3/internal/project"
	"github.com/stateful/runme/v3/internal/shell"
	"github.com/stateful/runme/v3/internal/term"
)

func listCmd(*commonFlags) *cobra.Command {
	var format string

	cmd := cobra.Command{
		Use:     "list [command1 command2 ...]",
		Aliases: []string{"ls"},
		Short:   "List commands.",
		Long: `List commands by optionally providing their names delimited by space.
The names are interpreted as glob patterns.

The --category option additionally filters the list of tasks to execute by category.`,
		Example: `List all blocks starting with the "generate-" prefix:
  runme beta list "generate-*"

List all blocks from the "setup" and "teardown" categories:
  runme beta list --category=setup,teardown
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return autoconfig.Invoke(
				func(
					proj *project.Project,
					filters []project.Filter,
					logger *zap.Logger,
					session *command.Session,
				) error {
					defer logger.Sync()

					tasks, err := project.LoadTasks(cmd.Context(), proj)
					if err != nil {
						return err
					}
					logger.Info("found tasks", zap.Int("count", len(tasks)))

					argsFilter, err := createProjectFilterFromPatterns(args)
					if err != nil {
						return err
					}

					filters = append(filters, argsFilter)

					tasks, err = project.FilterTasksByFn(tasks, filters...)
					if err != nil {
						return err
					}
					logger.Info("filtered tasks by filters", zap.Int("count", len(tasks)))

					switch format {
					case "json":
						return renderTasksAsJSONForCmd(cmd, tasks)
					case "table":
						return renderTasksAsTableForCmd(cmd, tasks)
					default:
						return errors.Errorf("invalid format: %s", format)
					}
				},
			)
		},
	}

	cmd.Flags().StringVar(&format, "format", "table", "Output format (table, json)")

	return &cmd
}

func renderTasksAsTableForCmd(cmd *cobra.Command, tasks []project.Task) error {
	term := term.FromIO(cmd.InOrStdin(), cmd.OutOrStdout(), cmd.ErrOrStderr())

	// Detect width. For non-TTY, use a default width of 80.
	width, _, err := term.Size()
	if err != nil {
		width = 80
	}

	table := tableprinter.New(term.Out(), term.IsTTY(), width)

	// table header
	table.AddField(strings.ToUpper("Name"))
	table.AddField(strings.ToUpper("File"))
	table.AddField(strings.ToUpper("First Command"))
	table.AddField(strings.ToUpper("Description"))
	table.AddField(strings.ToUpper("Named"))
	table.EndRow()

	for _, task := range tasks {
		named := "Yes"
		if task.CodeBlock.IsUnnamed() {
			named = "No"
		}

		table.AddField(task.CodeBlock.Name())
		table.AddField(relativePathToCwd(task.DocumentPath))
		table.AddField(renderLineFromLines(task.CodeBlock.Lines()))
		table.AddField(task.CodeBlock.Intro())
		table.AddField(named)
		table.EndRow()
	}

	return errors.WithStack(table.Render())
}

// TODO(adamb): output should be well-defined. It's questionable whether
// using [project.Task] is a good idea, as it comes from a different domain/layer.
// A person changing the [project.Task] structure may not be aware of the consequences.
func renderTasksAsJSONForCmd(cmd *cobra.Command, tasks []project.Task) error {
	raw, err := json.Marshal(tasks)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(
		jsonpretty.Format(cmd.OutOrStdout(), bytes.NewReader(raw), "  ", false),
	)
}

func renderLineFromLines(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	return shell.TryGetNonCommentLine(lines)
}
