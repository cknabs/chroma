package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// PowerQuery lexer.
var PowerQuery = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "PowerQuery",
		Aliases:         []string{"powerquery", "pq"},
		Filenames:       []string{"*.pq"},
		MimeTypes:       []string{"text/x-powerquery"},
		DotAll:          true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/powerquery.xml",
))
