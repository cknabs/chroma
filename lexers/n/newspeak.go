package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Newspeak lexer.
var Newspeak = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Newspeak",
		Aliases:   []string{"newspeak"},
		Filenames: []string{"*.ns2"},
		MimeTypes: []string{"text/x-newspeak"},
	},
	embedded,
	"embedded/newspeak.xml",
))
