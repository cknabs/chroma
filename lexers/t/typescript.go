package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TypeScript lexer.
var TypeScript = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "TypeScript",
		Aliases:   []string{"ts", "tsx", "typescript"},
		Filenames: []string{"*.ts", "*.tsx"},
		MimeTypes: []string{"text/x-typescript"},
		DotAll:    true,
		EnsureNL:  true,
	},
	embedded,
	"embedded/typescript.xml",
))
