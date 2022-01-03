package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ini lexer.
var Ini = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "INI",
		Aliases:   []string{"ini", "cfg", "dosini"},
		Filenames: []string{"*.ini", "*.cfg", "*.inf", ".gitconfig", ".editorconfig"},
		MimeTypes: []string{"text/x-ini", "text/inf"},
	},
	embedded,
	"embedded/ini.xml",
))
