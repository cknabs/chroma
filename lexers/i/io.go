package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Io lexer.
var Io = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Io",
		Aliases:   []string{"io"},
		Filenames: []string{"*.io"},
		MimeTypes: []string{"text/x-iosrc"},
	},
	embedded,
	"embedded/io.xml",
))
