package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Hy lexer.
var Hy = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Hy",
		Aliases:   []string{"hylang"},
		Filenames: []string{"*.hy"},
		MimeTypes: []string{"text/x-hy", "application/x-hy"},
	},
	embedded,
	"embedded/hy.xml",
))
