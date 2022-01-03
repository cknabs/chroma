package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// VB.Net lexer.
var VBNet = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "VB.net",
		Aliases:         []string{"vb.net", "vbnet"},
		Filenames:       []string{"*.vb", "*.bas"},
		MimeTypes:       []string{"text/x-vbnet", "text/x-vba"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/vb_net.xml",
))
