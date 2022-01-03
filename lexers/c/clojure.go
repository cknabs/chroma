package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Clojure lexer.
var Clojure = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Clojure",
		Aliases:   []string{"clojure", "clj"},
		Filenames: []string{"*.clj"},
		MimeTypes: []string{"text/x-clojure", "application/x-clojure"},
	},
	embedded,
	"embedded/clojure.xml",
))
