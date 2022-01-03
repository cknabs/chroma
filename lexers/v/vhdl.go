package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// VHDL lexer.
var VHDL = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "VHDL",
		Aliases:         []string{"vhdl"},
		Filenames:       []string{"*.vhdl", "*.vhd"},
		MimeTypes:       []string{"text/x-vhdl"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/vhdl.xml",
))
