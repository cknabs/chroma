package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Idris lexer.
var Idris = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Idris",
		Aliases:   []string{"idris", "idr"},
		Filenames: []string{"*.idr"},
		MimeTypes: []string{"text/x-idris"},
	},
	embedded,
	"embedded/idris.xml",
))
