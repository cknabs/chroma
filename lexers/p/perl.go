package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Perl lexer.
var Perl = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Perl",
		Aliases:   []string{"perl", "pl"},
		Filenames: []string{"*.pl", "*.pm", "*.t"},
		MimeTypes: []string{"text/x-perl", "application/x-perl"},
		DotAll:    true,
	},
	embedded,
	"embedded/perl.xml",
))
