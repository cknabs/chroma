package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Haskell lexer.
var Haskell = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Haskell",
		Aliases:   []string{"haskell", "hs"},
		Filenames: []string{"*.hs"},
		MimeTypes: []string{"text/x-haskell"},
	},
	embedded,
	"embedded/haskell.xml",
))
