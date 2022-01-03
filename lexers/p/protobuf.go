package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ProtocolBuffer lexer.
var ProtocolBuffer = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Protocol Buffer",
		Aliases:   []string{"protobuf", "proto"},
		Filenames: []string{"*.proto"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/protobuf.xml",
))
