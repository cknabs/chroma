package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TableGen lexer.
var Tablegen = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "TableGen",
		Aliases:   []string{"tablegen"},
		Filenames: []string{"*.td"},
		MimeTypes: []string{"text/x-tablegen"},
	},
	embedded,
	"embedded/tablegen.xml",
))
