package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Nginx Configuration File lexer.
var Nginx = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Nginx configuration file",
		Aliases:   []string{"nginx"},
		Filenames: []string{"nginx.conf"},
		MimeTypes: []string{"text/x-nginx-conf"},
	},
	embedded,
	"embedded/nginx.xml",
))
