package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Java lexer.
var Java = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Java",
		Aliases:   []string{"java"},
		Filenames: []string{"*.java"},
		MimeTypes: []string{"text/x-java"},
		DotAll:    true,
		EnsureNL:  true,
	},
	embedded,
	"embedded/java.xml",
))
