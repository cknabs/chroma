package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Tasm lexer.
var Tasm = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "TASM",
		Aliases:         []string{"tasm"},
		Filenames:       []string{"*.asm", "*.ASM", "*.tasm"},
		MimeTypes:       []string{"text/x-tasm"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/tasm.xml",
))
