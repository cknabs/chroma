package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// C lexer.
var C = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "C",
		Aliases:   []string{"c"},
		Filenames: []string{"*.c", "*.h", "*.idc", "*.x[bp]m"},
		MimeTypes: []string{"text/x-chdr", "text/x-csrc", "image/x-xbitmap", "image/x-xpixmap"},
		EnsureNL:  true,
	},
	embedded,
	"embedded/c.xml",
))
