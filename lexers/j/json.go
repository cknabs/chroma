package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// JSON lexer.
var JSON = internal.Register(MustNewXMLLexer(
	&Config{
		Name:         "JSON",
		Aliases:      []string{"json"},
		Filenames:    []string{"*.json"},
		MimeTypes:    []string{"application/json"},
		NotMultiline: true,
		DotAll:       true,
	},
	embedded,
	"embedded/json.xml",
))
