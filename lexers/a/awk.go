package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Awk lexer.
var Awk = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Awk",
		Aliases:   []string{"awk", "gawk", "mawk", "nawk"},
		Filenames: []string{"*.awk"},
		MimeTypes: []string{"application/x-awk"},
	},
	embedded,
	"embedded/awk.xml",
))
