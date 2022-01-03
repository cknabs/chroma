package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cfstatement lexer.
var Cfstatement = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "cfstatement",
		Aliases:         []string{"cfs"},
		Filenames:       []string{},
		MimeTypes:       []string{},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/cfengine3.xml",
))
