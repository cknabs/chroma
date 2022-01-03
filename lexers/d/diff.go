package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Diff lexer.
var Diff = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Diff",
		Aliases:   []string{"diff", "udiff"},
		EnsureNL:  true,
		Filenames: []string{"*.diff", "*.patch"},
		MimeTypes: []string{"text/x-diff", "text/x-patch"},
	},
	embedded,
	"embedded/diff.xml",
))
