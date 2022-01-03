package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Verilog lexer.
var Verilog = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "verilog",
		Aliases:   []string{"verilog", "v"},
		Filenames: []string{"*.v"},
		MimeTypes: []string{"text/x-verilog"},
		EnsureNL:  true,
	},
	embedded,
	"embedded/verilog.xml",
))
