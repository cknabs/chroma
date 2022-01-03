package k

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Kotlin lexer.
var Kotlin = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Kotlin",
		Aliases:   []string{"kotlin"},
		Filenames: []string{"*.kt"},
		MimeTypes: []string{"text/x-kotlin"},
		DotAll:    true,
	},
	embedded,
	"embedded/kotlin.xml",
))
