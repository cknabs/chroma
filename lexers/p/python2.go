package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Python2 lexer.
var Python2 = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Python 2",
		Aliases:   []string{"python2", "py2"},
		Filenames: []string{},
		MimeTypes: []string{"text/x-python2", "application/x-python2"},
	},
	embedded,
	"embedded/python_2.xml",
))
