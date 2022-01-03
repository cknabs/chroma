package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MLIR lexer.
var Mlir = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "MLIR",
		Aliases:   []string{"mlir"},
		Filenames: []string{"*.mlir"},
		MimeTypes: []string{"text/x-mlir"},
	},
	embedded,
	"embedded/mlir.xml",
))
