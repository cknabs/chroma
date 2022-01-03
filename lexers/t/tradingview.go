package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TradingView lexer
var TradingView = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "TradingView",
		Aliases:   []string{"tradingview", "tv"},
		Filenames: []string{"*.tv"},
		MimeTypes: []string{"text/x-tradingview"},
		DotAll:    true,
		EnsureNL:  true,
	},
	embedded,
	"embedded/tradingview.xml",
))
