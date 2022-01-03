package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Squidconf lexer.
var Squidconf = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "SquidConf",
		Aliases:         []string{"squidconf", "squid.conf", "squid"},
		Filenames:       []string{"squid.conf"},
		MimeTypes:       []string{"text/x-squidconf"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/squidconf.xml",
))
