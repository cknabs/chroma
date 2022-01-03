package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Applescript lexer.
var Applescript = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "AppleScript",
		Aliases:   []string{"applescript"},
		Filenames: []string{"*.applescript"},
		MimeTypes: []string{},
		DotAll:    true,
	},
	embedded,
	"embedded/applescript.xml",
))
