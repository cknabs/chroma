package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Metal lexer.
var Metal = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Metal",
		Aliases:   []string{"metal"},
		Filenames: []string{"*.metal"},
		MimeTypes: []string{"text/x-metal"},
		EnsureNL:  true,
	},
	embedded,
	"embedded/metal.xml",
))
