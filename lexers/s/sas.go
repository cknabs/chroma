package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Sas lexer.
var Sas = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "SAS",
		Aliases:         []string{"sas"},
		Filenames:       []string{"*.SAS", "*.sas"},
		MimeTypes:       []string{"text/x-sas", "text/sas", "application/x-sas"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/sas.xml",
))
