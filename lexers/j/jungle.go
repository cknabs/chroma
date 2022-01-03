package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var Jungle = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Jungle",
		Aliases:   []string{"jungle"},
		Filenames: []string{"*.jungle"},
		MimeTypes: []string{"text/x-jungle"},
	},
	embedded,
	"embedded/jungle.xml",
))
