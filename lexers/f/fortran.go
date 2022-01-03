package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Fortran lexer.
var Fortran = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Fortran",
		Aliases:         []string{"fortran"},
		Filenames:       []string{"*.f03", "*.f90", "*.F03", "*.F90"},
		MimeTypes:       []string{"text/x-fortran"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/fortran.xml",
))
