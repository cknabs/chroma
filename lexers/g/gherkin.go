package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Gherkin lexer.
var Gherkin = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Gherkin",
		Aliases:   []string{"cucumber", "Cucumber", "gherkin", "Gherkin"},
		Filenames: []string{"*.feature", "*.FEATURE"},
		MimeTypes: []string{"text/x-gherkin"},
	},
	embedded,
	"embedded/gherkin.xml",
))
