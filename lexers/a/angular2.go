package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Angular2 lexer.
var Angular2 = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Angular2",
		Aliases:   []string{"ng2"},
		Filenames: []string{},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/angular2.xml",
))
