package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ada lexer.
var Ada = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Ada",
		Aliases:         []string{"ada", "ada95", "ada2005"},
		Filenames:       []string{"*.adb", "*.ads", "*.ada"},
		MimeTypes:       []string{"text/x-ada"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/ada.xml",
))
