package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Tcl lexer.
var Tcl = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Tcl",
		Aliases:   []string{"tcl"},
		Filenames: []string{"*.tcl", "*.rvt"},
		MimeTypes: []string{"text/x-tcl", "text/x-script.tcl", "application/x-tcl"},
	},
	embedded,
	"embedded/tcl.xml",
))
