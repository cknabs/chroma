package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// nolint

// Lexer for the Plutus Core Languages (version 2.1)
//
// including both Typed- and Untyped- versions
// based on “Formal Specification of the Plutus Core Language (version 2.1)”, published 6th April 2021:
// https://hydra.iohk.io/build/8205579/download/1/plutus-core-specification.pdf

var PlutusCoreLang = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Plutus Core",
		Aliases:   []string{"plutus-core", "plc"},
		Filenames: []string{"*.plc"},
		MimeTypes: []string{"text/x-plutus-core", "application/x-plutus-core"},
	},
	embedded,
	"embedded/plutus_core.xml",
))
