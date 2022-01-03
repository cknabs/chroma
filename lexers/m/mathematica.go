package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Mathematica lexer.
var Mathematica = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Mathematica",
		Aliases:   []string{"mathematica", "mma", "nb"},
		Filenames: []string{"*.nb", "*.cdf", "*.nbp", "*.ma"},
		MimeTypes: []string{"application/mathematica", "application/vnd.wolfram.mathematica", "application/vnd.wolfram.mathematica.package", "application/vnd.wolfram.cdf"},
	},
	embedded,
	"embedded/mathematica.xml",
))
