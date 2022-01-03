package v

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Vue lexer.
//
// This was generated from https://github.com/testdrivenio/vue-lexer
var Vue = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "vue",
		Aliases:   []string{"vue", "vuejs"},
		Filenames: []string{"*.vue"},
		MimeTypes: []string{"text/x-vue", "application/x-vue"},
		DotAll:    true,
	},
	embedded,
	"embedded/vue.xml",
))
