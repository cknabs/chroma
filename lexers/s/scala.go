package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Scala lexer.
var Scala = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Scala",
		Aliases:   []string{"scala"},
		Filenames: []string{"*.scala"},
		MimeTypes: []string{"text/x-scala"},
		DotAll:    true,
	},
	embedded,
	"embedded/scala.xml",
))
