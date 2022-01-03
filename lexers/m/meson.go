package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Meson lexer.
var Meson = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Meson",
		Aliases:   []string{"meson", "meson.build"},
		Filenames: []string{"meson.build", "meson_options.txt"},
		MimeTypes: []string{"text/x-meson"},
	},
	embedded,
	"embedded/meson.xml",
))
