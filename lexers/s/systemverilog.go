package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Systemverilog lexer.
var Systemverilog = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "systemverilog",
		Aliases:   []string{"systemverilog", "sv"},
		Filenames: []string{"*.sv", "*.svh"},
		MimeTypes: []string{"text/x-systemverilog"},
		EnsureNL:  true,
	},
	embedded,
	"embedded/systemverilog.xml",
))
