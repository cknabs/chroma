package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Igor lexer.
var Igor = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Igor",
		Aliases:         []string{"igor", "igorpro"},
		Filenames:       []string{"*.ipf"},
		MimeTypes:       []string{"text/ipf"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/igor.xml",
))
