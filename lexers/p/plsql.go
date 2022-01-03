package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pl/Pgsql lexer.
var PLpgSQL = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "PL/pgSQL",
		Aliases:         []string{"plpgsql"},
		Filenames:       []string{},
		MimeTypes:       []string{"text/x-plpgsql"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/pl_pgsql.xml",
))
