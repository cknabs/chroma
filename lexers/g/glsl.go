package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// GLSL lexer.
var GLSL = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "GLSL",
		Aliases:   []string{"glsl"},
		Filenames: []string{"*.vert", "*.frag", "*.geo"},
		MimeTypes: []string{"text/x-glslsrc"},
	},
	embedded,
	"embedded/glsl.xml",
))
