package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Bibtex lexer.
var Bibtex = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "BibTeX",
		Aliases:         []string{"bib", "bibtex"},
		Filenames:       []string{"*.bib"},
		MimeTypes:       []string{"text/x-bibtex"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/bibtex.xml",
))
