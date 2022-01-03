package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Reg lexer.
var Reg = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "reg",
		Aliases:   []string{"registry"},
		Filenames: []string{"*.reg"},
		MimeTypes: []string{"text/x-windows-registry"},
	},
	embedded,
	"embedded/reg.xml",
))
