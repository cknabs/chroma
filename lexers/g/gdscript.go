package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// GDScript lexer.
var GDScript = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "GDScript",
		Aliases:   []string{"gdscript", "gd"},
		Filenames: []string{"*.gd"},
		MimeTypes: []string{"text/x-gdscript", "application/x-gdscript"},
	},
	embedded,
	"embedded/gdscript.xml",
))
