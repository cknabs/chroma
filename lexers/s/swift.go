package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Swift lexer.
var Swift = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Swift",
		Aliases:   []string{"swift"},
		Filenames: []string{"*.swift"},
		MimeTypes: []string{"text/x-swift"},
	},
	embedded,
	"embedded/swift.xml",
))
