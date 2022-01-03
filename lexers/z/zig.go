package z

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Zig lexer.
var Zig = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Zig",
		Aliases:   []string{"zig"},
		Filenames: []string{"*.zig"},
		MimeTypes: []string{"text/zig"},
	},
	embedded,
	"embedded/zig.xml",
))
