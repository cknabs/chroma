package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Django/Jinja lexer.
var DjangoJinja = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Django/Jinja",
		Aliases:   []string{"django", "jinja"},
		Filenames: []string{},
		MimeTypes: []string{"application/x-django-templating", "application/x-jinja"},
		DotAll:    true,
	},
	embedded,
	"embedded/django_jinja.xml",
))
