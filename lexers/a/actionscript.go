package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Actionscript lexer.
var Actionscript = internal.Register(MustNewXMLLexer(
	&Config{
		Name:         "ActionScript",
		Aliases:      []string{"as", "actionscript"},
		Filenames:    []string{"*.as"},
		MimeTypes:    []string{"application/x-actionscript", "text/x-actionscript", "text/actionscript"},
		NotMultiline: true,
		DotAll:       true,
	},
	embedded,
	"embedded/actionscript.xml",
))
