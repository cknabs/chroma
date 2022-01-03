package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Al lexer.
var Al = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "AL",
		Aliases:         []string{"al"},
		Filenames:       []string{"*.al", "*.dal"},
		MimeTypes:       []string{"text/x-al"},
		DotAll:          true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/al.xml",
))
