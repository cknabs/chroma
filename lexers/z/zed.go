package z

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Zed lexer.
var Zed = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Zed",
		Aliases:   []string{"zed"},
		Filenames: []string{"*.zed"},
		MimeTypes: []string{"text/zed"},
	},
	embedded,
	"embedded/zed.xml",
).SetAnalyser(func(text string) float32 {
	if strings.Contains(text, "definition ") && strings.Contains(text, "relation ") && strings.Contains(text, "permission ") {
		return 0.9
	}
	if strings.Contains(text, "definition ") {
		return 0.5
	}
	if strings.Contains(text, "relation ") {
		return 0.5
	}
	if strings.Contains(text, "permission ") {
		return 0.25
	}
	return 0.0
}))
