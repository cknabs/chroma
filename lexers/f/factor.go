package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Factor lexer.
var Factor = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Factor",
		Aliases:   []string{"factor"},
		Filenames: []string{"*.factor"},
		MimeTypes: []string{"text/x-factor"},
	},
	embedded,
	"embedded/factor.xml",
))
