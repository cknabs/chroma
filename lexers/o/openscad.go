package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var OpenSCAD = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "OpenSCAD",
		Aliases:   []string{"openscad"},
		Filenames: []string{"*.scad"},
		MimeTypes: []string{"text/x-scad"},
	},
	embedded,
	"embedded/openscad.xml",
))
