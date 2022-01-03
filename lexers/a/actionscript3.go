package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Actionscript 3 lexer.
var Actionscript3 = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "ActionScript 3",
		Aliases:   []string{"as3", "actionscript3"},
		Filenames: []string{"*.as"},
		MimeTypes: []string{"application/x-actionscript3", "text/x-actionscript3", "text/actionscript3"},
		DotAll:    true,
	},
	embedded,
	"embedded/actionscript_3.xml",
))
