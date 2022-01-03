package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Dart lexer.
var Dart = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Dart",
		Aliases:   []string{"dart"},
		Filenames: []string{"*.dart"},
		MimeTypes: []string{"text/x-dart"},
		DotAll:    true,
	},
	embedded,
	"embedded/dart.xml",
))
