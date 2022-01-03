package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// mcfunction lexer.
var MCFunction = internal.Register(MustNewXMLLexer(
	&Config{
		Name:         "mcfunction",
		Aliases:      []string{"mcfunction"},
		Filenames:    []string{"*.mcfunction"},
		MimeTypes:    []string{},
		NotMultiline: true,
		DotAll:       true,
	},
	embedded,
	"embedded/mcfunction.xml",
))
