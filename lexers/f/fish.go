package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Fish lexer.
var Fish = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Fish",
		Aliases:   []string{"fish", "fishshell"},
		Filenames: []string{"*.fish", "*.load"},
		MimeTypes: []string{"application/x-fish"},
	},
	embedded,
	"embedded/fish.xml",
))
