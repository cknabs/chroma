package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Nix lexer.
var Nix = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Nix",
		Aliases:   []string{"nixos", "nix"},
		Filenames: []string{"*.nix"},
		MimeTypes: []string{"text/x-nix"},
	},
	embedded,
	"embedded/nix.xml",
))
