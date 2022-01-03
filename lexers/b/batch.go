package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Batchfile lexer.
var Batchfile = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Batchfile",
		Aliases:         []string{"bat", "batch", "dosbatch", "winbatch"},
		Filenames:       []string{"*.bat", "*.cmd"},
		MimeTypes:       []string{"application/x-dos-batch"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/batchfile.xml",
))
