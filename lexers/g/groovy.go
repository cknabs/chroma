package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Groovy lexer.
var Groovy = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Groovy",
		Aliases:   []string{"groovy"},
		Filenames: []string{"*.groovy", "*.gradle"},
		MimeTypes: []string{"text/x-groovy"},
		DotAll:    true,
	},
	embedded,
	"embedded/groovy.xml",
))
