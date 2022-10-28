package document

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCodeBlock(t *testing.T) {
	data := []byte(`
This is a basic snippet with a shell command:

` + "```" + `sh
$ echo "Hello, runme!"
` + "```" + `

It can have an annotation with a name:

` + "```" + `sh {name=echo first= second=2}
$ echo "Hello, runme!"
` + "```")
	source := NewSource(data)
	blocks := source.Parse().CodeBlocks()

	assert.EqualValues(t, map[string]string{}, blocks[0].Attributes())
	assert.Equal(t, "echo-hello", blocks[0].Name())
	assert.Equal(t, "```sh\n$ echo \"Hello, runme!\"\n```", blocks[0].Content())

	assert.EqualValues(t, map[string]string{
		"name":   "echo",
		"first":  "",
		"second": "2",
	}, blocks[1].Attributes())
	assert.Equal(t, "echo", blocks[1].Name())
	assert.Equal(t, "```sh {name=echo first= second=2}\n$ echo \"Hello, runme!\"\n```", blocks[1].Content())
	assert.Equal(t, []string{`echo "Hello, runme!"`}, blocks[1].Lines())
	assert.Equal(t, "sh", blocks[1].Executable())
	assert.Equal(t, "It can have an annotation with a name.", blocks[1].Intro())
}

func TestCodeBlock_MapLines(t *testing.T) {
	data := []byte("```" + `sh
echo 1
echo 2
echo 3
` + "```")
	source := NewSource(data)
	block := source.Parse().CodeBlocks()[0]
	err := block.MapLines(func(s string) (string, error) {
		return strings.Split(s, " ")[1], nil
	})
	require.NoError(t, err)
	assert.Equal(t, []string{"1", "2", "3"}, block.Lines())
}

func TestMarkdownBlock_Content(t *testing.T) {
	data := []byte(`This is a basic snippet with a shell command:

` + "```" + `sh
$ echo "Hello, runme!"
` + "```" + `

It can have an annotation with a name:

` + "```" + `sh {name=myname first= second=2}
$ echo "Hello, runme!"
` + "```")
	source := NewSource(data)
	blocks := source.Parse().Blocks()
	assert.Equal(t, "This is a basic snippet with a shell command:", blocks[0].(*MarkdownBlock).Content())
	assert.Equal(t, "It can have an annotation with a name:", blocks[2].(*MarkdownBlock).Content())
}

func TestMarkdownBlock_Content2(t *testing.T) {
	data := []byte(`  > bq 1
  > bq 2

1. Item 1
1. Item 2

` + "```sh" + `
echo 1
echo 2
` + "```" + `

   a
   b
`)
	source := NewSource(data)
	blocks := source.Parse().Blocks()
	assert.Equal(t, 4, len(blocks))
	assert.Equal(t, "  > bq 1\n  > bq 2", blocks[0].(*MarkdownBlock).Content())
	assert.Equal(t, "1. Item 1\n1. Item 2", blocks[1].(*MarkdownBlock).Content())
	assert.Equal(t, "```sh\necho 1\necho 2\n```", blocks[2].(*CodeBlock).Content())
	assert.Equal(t, "   a\n   b", blocks[3].(*MarkdownBlock).Content())
}
