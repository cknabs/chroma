package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var TOML = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "TOML",
		Aliases:   []string{"toml"},
		Filenames: []string{"*.toml"},
		MimeTypes: []string{"text/x-toml"},
	},
	embedded,
	"embedded/toml.xml",
))
