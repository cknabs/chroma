package y

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var YANG = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "YANG",
		Aliases:   []string{"yang"},
		Filenames: []string{"*.yang"},
		MimeTypes: []string{"application/yang"},
	},
	embedded,
	"embedded/yang.xml",
))
