package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// D lexer. https://dlang.org/spec/lex.html
var D = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "D",
		Aliases:   []string{"d"},
		Filenames: []string{"*.d", "*.di"},
		MimeTypes: []string{"text/x-d"},
		EnsureNL:  true,
	},
	embedded,
	"embedded/d.xml",
))
