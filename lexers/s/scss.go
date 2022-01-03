package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Scss lexer.
var Scss = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "SCSS",
		Aliases:         []string{"scss"},
		Filenames:       []string{"*.scss"},
		MimeTypes:       []string{"text/x-scss"},
		NotMultiline:    true,
		DotAll:          true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/scss.xml",
))
