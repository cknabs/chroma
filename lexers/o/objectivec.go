package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Objective-C lexer.
var ObjectiveC = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Objective-C",
		Aliases:   []string{"objective-c", "objectivec", "obj-c", "objc"},
		Filenames: []string{"*.m", "*.h"},
		MimeTypes: []string{"text/x-objective-c"},
	},
	embedded,
	"embedded/objective-c.xml",
))
