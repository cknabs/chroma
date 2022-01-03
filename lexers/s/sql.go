package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// SQL lexer.
var SQL = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "SQL",
		Aliases:         []string{"sql"},
		Filenames:       []string{"*.sql"},
		MimeTypes:       []string{"text/x-sql"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/sql.xml",
))
