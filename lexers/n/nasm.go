package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Nasm lexer.
var Nasm = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "NASM",
		Aliases:         []string{"nasm"},
		Filenames:       []string{"*.asm", "*.ASM"},
		MimeTypes:       []string{"text/x-nasm"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/nasm.xml",
))
