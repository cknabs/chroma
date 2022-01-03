package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Puppet lexer.
var Puppet = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Puppet",
		Aliases:   []string{"puppet"},
		Filenames: []string{"*.pp"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/puppet.xml",
))
