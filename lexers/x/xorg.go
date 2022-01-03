package x

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Xorg lexer.
var Xorg = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Xorg",
		Aliases:   []string{"xorg.conf"},
		Filenames: []string{"xorg.conf"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/xorg.xml",
))
