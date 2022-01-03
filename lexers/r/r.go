package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// R/S lexer.
var R = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "R",
		Aliases:   []string{"splus", "s", "r"},
		Filenames: []string{"*.S", "*.R", "*.r", ".Rhistory", ".Rprofile", ".Renviron"},
		MimeTypes: []string{"text/S-plus", "text/S", "text/x-r-source", "text/x-r", "text/x-R", "text/x-r-history", "text/x-r-profile"},
	},
	embedded,
	"embedded/r.xml",
))
