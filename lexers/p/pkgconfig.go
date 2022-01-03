package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pkgconfig lexer.
var Pkgconfig = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "PkgConfig",
		Aliases:   []string{"pkgconfig"},
		Filenames: []string{"*.pc"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/pkgconfig.xml",
))
