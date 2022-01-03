package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Tex lexer.
var TeX = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "TeX",
		Aliases:   []string{"tex", "latex"},
		Filenames: []string{"*.tex", "*.aux", "*.toc"},
		MimeTypes: []string{"text/x-tex", "text/x-latex"},
	},
	embedded,
	"embedded/tex.xml",
))
