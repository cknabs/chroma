package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TransactSQL lexer.
var TransactSQL = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Transact-SQL",
		Aliases:         []string{"tsql", "t-sql"},
		MimeTypes:       []string{"text/x-tsql"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/transact-sql.xml",
))
