package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MorrowindScript lexer.
var MorrowindScript = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "MorrowindScript",
		Aliases:   []string{"morrowind", "mwscript"},
		Filenames: []string{},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/morrowindscript.xml",
))
