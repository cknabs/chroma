package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Thrift lexer.
var Thrift = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Thrift",
		Aliases:   []string{"thrift"},
		Filenames: []string{"*.thrift"},
		MimeTypes: []string{"application/x-thrift"},
	},
	embedded,
	"embedded/thrift.xml",
))
