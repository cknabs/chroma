package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cfengine3 lexer.
var Cfengine3 = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "CFEngine3",
		Aliases:   []string{"cfengine3", "cf3"},
		Filenames: []string{"*.cf"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/cfengine3.xml",
))
