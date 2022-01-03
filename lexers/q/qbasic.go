package q

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Qbasic lexer.
var Qbasic = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "QBasic",
		Aliases:   []string{"qbasic", "basic"},
		Filenames: []string{"*.BAS", "*.bas"},
		MimeTypes: []string{"text/basic"},
	},
	embedded,
	"embedded/qbasic.xml",
))
