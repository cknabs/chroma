package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Llvm lexer.
var Llvm = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "LLVM",
		Aliases:   []string{"llvm"},
		Filenames: []string{"*.ll"},
		MimeTypes: []string{"text/x-llvm"},
	},
	embedded,
	"embedded/llvm.xml",
))
