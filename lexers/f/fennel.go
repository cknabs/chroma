package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Fennel lexer.
var Fennel = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Fennel",
		Aliases:   []string{"fennel", "fnl"},
		Filenames: []string{"*.fennel"},
		MimeTypes: []string{"text/x-fennel", "application/x-fennel"},
	},
	embedded,
	"embedded/fennel.xml",
))

// Here's some Fennel code used to generate the lists of keywords:
// (local fennel (require :fennel))
//
// (fn member? [t x] (each [_ y (ipairs t)] (when (= y x) (lua "return true"))))
//
// (local declarations [:fn :lambda :λ :local :var :global :macro :macros])
// (local keywords [])
// (local globals [])
//
// (each [name data (pairs (fennel.syntax))]
//   (if (member? declarations name) nil ; already populated
//       data.special? (table.insert keywords name)
//       data.macro? (table.insert keywords name)
//       data.global? (table.insert globals name)))
//
// (fn quoted [tbl]
//   (table.sort tbl)
//   (table.concat (icollect [_ k (ipairs tbl)]
//                   (string.format "`%s`" k)) ", "))
//
// (print :Keyword (quoted keywords))
// (print :KeywordDeclaration (quoted declarations))
// (print :NameBuiltin (quoted globals))
