package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cython lexer.
var Cython = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Cython",
		Aliases:   []string{"cython", "pyx", "pyrex"},
		Filenames: []string{"*.pyx", "*.pxd", "*.pxi"},
		MimeTypes: []string{"text/x-cython", "application/x-cython"},
	},
	embedded,
	"embedded/cython.xml",
))
