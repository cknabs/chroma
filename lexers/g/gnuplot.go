package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Gnuplot lexer.
var Gnuplot = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Gnuplot",
		Aliases:   []string{"gnuplot"},
		Filenames: []string{"*.plot", "*.plt"},
		MimeTypes: []string{"text/x-gnuplot"},
	},
	embedded,
	"embedded/gnuplot.xml",
))
