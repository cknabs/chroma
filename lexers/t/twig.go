package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Twig lexer.
var Twig = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Twig",
		Aliases:   []string{"twig"},
		Filenames: []string{},
		MimeTypes: []string{"application/x-twig"},
		DotAll:    true,
	},
	embedded,
	"embedded/twig.xml",
))
