package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ruby lexer.
var Ruby = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Ruby",
		Aliases:   []string{"rb", "ruby", "duby"},
		Filenames: []string{"*.rb", "*.rbw", "Rakefile", "*.rake", "*.gemspec", "*.rbx", "*.duby", "Gemfile"},
		MimeTypes: []string{"text/x-ruby", "application/x-ruby"},
		DotAll:    true,
	},
	embedded,
	"embedded/ruby.xml",
))
