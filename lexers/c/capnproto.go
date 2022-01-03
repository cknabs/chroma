package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cap'N'Proto Proto lexer.
var CapNProto = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Cap'n Proto",
		Aliases:   []string{"capnp"},
		Filenames: []string{"*.capnp"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/cap_n_proto.xml",
))
