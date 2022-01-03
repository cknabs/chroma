package circular

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// PHP lexer for pure PHP code (not embedded in HTML).
var PHP = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "PHP",
		Aliases:         []string{"php", "php3", "php4", "php5"},
		Filenames:       []string{"*.php", "*.php[345]", "*.inc"},
		MimeTypes:       []string{"text/x-php"},
		DotAll:          true,
		CaseInsensitive: true,
		EnsureNL:        true,
	},
	embedded,
	"embedded/php.xml",
))
