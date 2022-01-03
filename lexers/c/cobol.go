package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cobol lexer.
var Cobol = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "COBOL",
		Aliases:         []string{"cobol"},
		Filenames:       []string{"*.cob", "*.COB", "*.cpy", "*.CPY"},
		MimeTypes:       []string{"text/x-cobol"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/cobol.xml",
))
