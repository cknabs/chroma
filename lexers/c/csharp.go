package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CSharp lexer.
var CSharp = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "C#",
		Aliases:   []string{"csharp", "c#"},
		Filenames: []string{"*.cs"},
		MimeTypes: []string{"text/x-csharp"},
		DotAll:    true,
		EnsureNL:  true,
	},
	embedded,
	"embedded/csharp.xml",
))
