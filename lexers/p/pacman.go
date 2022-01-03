package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pacmanconf lexer.
var Pacmanconf = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "PacmanConf",
		Aliases:   []string{"pacmanconf"},
		Filenames: []string{"pacman.conf"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/pacmanconf.xml",
))
