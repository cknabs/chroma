package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Racket lexer.
var Racket = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Racket",
		Aliases:   []string{"racket", "rkt"},
		Filenames: []string{"*.rkt", "*.rktd", "*.rktl"},
		MimeTypes: []string{"text/x-racket", "application/x-racket"},
	},
	embedded,
	"embedded/racket.xml",
))
