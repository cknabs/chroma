package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var MonkeyC = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "MonkeyC",
		Aliases:   []string{"monkeyc"},
		Filenames: []string{"*.mc"},
		MimeTypes: []string{"text/x-monkeyc"},
	},
	embedded,
	"embedded/monkeyc.xml",
))
