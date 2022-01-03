package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Gas lexer.
var Gas = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "GAS",
		Aliases:   []string{"gas", "asm"},
		Filenames: []string{"*.s", "*.S"},
		MimeTypes: []string{"text/x-gas"},
	},
	embedded,
	"embedded/gas.xml",
))
