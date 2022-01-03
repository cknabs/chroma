package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Scilab lexer.
var Scilab = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Scilab",
		Aliases:   []string{"scilab"},
		Filenames: []string{"*.sci", "*.sce", "*.tst"},
		MimeTypes: []string{"text/scilab"},
	},
	embedded,
	"embedded/scilab.xml",
))
