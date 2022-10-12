package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeBlock_Attributes(t *testing.T) {
	data := []byte(`
This is a basic snippet with a shell command:

` + "```" + `sh
$ echo "Hello, rdme!"
` + "```" + `

It can have an annotation with a name:

` + "```" + `sh {name=echo first= second=2}
$ echo "Hello, rdme!"
` + "```")
	source := NewSource(data)
	blocks := source.Parse().CodeBlocks()
	assert.EqualValues(t, map[string]string{}, blocks[0].Attributes())
	assert.EqualValues(t, map[string]string{
		"name":   "echo",
		"first":  "",
		"second": "2",
	}, blocks[1].Attributes())
}

func TestCodeBlock_Content(t *testing.T) {
	data := []byte("```" + `sh
$ echo "Hello, rdme!"
` + "```")
	source := NewSource(data)
	block := source.Parse().CodeBlocks()[0]
	assert.Equal(t, "$ echo \"Hello, rdme!\"\n", block.Content())

	block.SetContent("reset")
	assert.Equal(t, "reset", block.Content())
}

func TestCodeBlock_Executable(t *testing.T) {
	data := []byte("```" + `sh
$ echo "Hello, rdme!"
` + "```")
	source := NewSource(data)
	block := source.Parse().CodeBlocks()[0]
	assert.Equal(t, "sh", block.Executable())
}

func TestCodeBlock_Intro(t *testing.T) {
	data := []byte(`
This is a basic snippet with a shell command:

` + "```" + `sh
$ echo "Hello, rdme!"
` + "```")
	source := NewSource(data)
	block := source.Parse().CodeBlocks()[0]
	assert.Equal(t, "This is a basic snippet with a shell command.", block.Intro())
}

func TestCodeBlock_Lines(t *testing.T) {
	source := NewSource(testREADME)
	blocks := source.Parse().CodeBlocks()
	assert.Len(t, blocks, 5)
	assert.Equal(t, 1, blocks[0].LineCount())
	assert.Equal(t, "echo \"Hello, rdme!\"", blocks[0].Line(0))
	assert.Equal(t, []string{
		"echo \"1\"",
		"echo \"2\"",
		"echo \"3\"",
	}, blocks[2].Lines())
}

func TestCodeBlock_Name(t *testing.T) {
	data := []byte(`
This is a basic snippet with a shell command:

` + "```" + `sh
$ echo "Hello, rdme!"
` + "```" + `

It can have an annotation with a name:

` + "```" + `sh {name=myname first= second=2}
$ echo "Hello, rdme!"
` + "```")
	source := NewSource(data)
	blocks := source.Parse().CodeBlocks()
	assert.Equal(t, "echo-hello", blocks[0].Name())
	assert.Equal(t, "myname", blocks[1].Name())
}