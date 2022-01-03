package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Lua lexer.
var Lua = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Lua",
		Aliases:   []string{"lua"},
		Filenames: []string{"*.lua", "*.wlua"},
		MimeTypes: []string{"text/x-lua", "application/x-lua"},
	},
	embedded,
	"embedded/lua.xml",
))
