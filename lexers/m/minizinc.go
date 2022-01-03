package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// MiniZinc lexer.
var MZN = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "MiniZinc",
		Aliases:   []string{"minizinc", "MZN", "mzn"},
		Filenames: []string{"*.mzn", "*.dzn", "*.fzn"},
		MimeTypes: []string{"text/minizinc"},
	},
	embedded,
	"embedded/minizinc.xml",
))
