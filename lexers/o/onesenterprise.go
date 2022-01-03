package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// 1S:Enterprise lexer.
var OnesEnterprise = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "OnesEnterprise",
		Aliases:         []string{"ones", "onesenterprise", "1S", "1S:Enterprise"},
		Filenames:       []string{"*.EPF", "*.epf", "*.ERF", "*.erf"},
		MimeTypes:       []string{"application/octet-stream"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/onesenterprise.xml",
))
