package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Lighttpd Configuration File lexer.
var Lighttpd = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Lighttpd configuration file",
		Aliases:   []string{"lighty", "lighttpd"},
		Filenames: []string{},
		MimeTypes: []string{"text/x-lighttpd-conf"},
	},
	embedded,
	"embedded/lighttpd.xml",
))
