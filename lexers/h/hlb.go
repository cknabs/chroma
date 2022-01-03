package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// HLB lexer.
var HLB = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "HLB",
		Aliases:   []string{"hlb"},
		Filenames: []string{"*.hlb"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/hlb.xml",
))
