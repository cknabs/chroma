package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Apacheconf lexer.
var Apacheconf = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "ApacheConf",
		Aliases:         []string{"apacheconf", "aconf", "apache"},
		Filenames:       []string{".htaccess", "apache.conf", "apache2.conf"},
		MimeTypes:       []string{"text/x-apacheconf"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/apacheconf.xml",
))
