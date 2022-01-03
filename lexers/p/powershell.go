package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Powershell lexer.
var Powershell = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "PowerShell",
		Aliases:         []string{"powershell", "posh", "ps1", "psm1", "psd1"},
		Filenames:       []string{"*.ps1", "*.psm1", "*.psd1"},
		MimeTypes:       []string{"text/x-powershell"},
		DotAll:          true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/powershell.xml",
))
