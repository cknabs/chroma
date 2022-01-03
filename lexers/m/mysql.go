package m

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	mysqlAnalyserNameBetweenBacktickRe = regexp.MustCompile("`[a-zA-Z_]\\w*`")
	mysqlAnalyserNameBetweenBracketRe  = regexp.MustCompile(`\[[a-zA-Z_]\w*\]`)
)

// MySQL lexer.
var MySQL = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "MySQL",
		Aliases:         []string{"mysql"},
		Filenames:       []string{"*.sql"},
		MimeTypes:       []string{"text/x-mysql"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/mysql.xml",
).SetAnalyser(func(text string) float32 {
	nameBetweenBacktickCount := len(mysqlAnalyserNameBetweenBacktickRe.FindAllString(text, -1))
	nameBetweenBracketCount := len(mysqlAnalyserNameBetweenBracketRe.FindAllString(text, -1))

	var result float32

	// Same logic as above in the TSQL analysis.
	dialectNameCount := nameBetweenBacktickCount + nameBetweenBracketCount
	if dialectNameCount >= 1 && nameBetweenBacktickCount >= (2*nameBetweenBracketCount) {
		// Found at least twice as many `name` as [name].
		result += 0.5
	} else if nameBetweenBacktickCount > nameBetweenBracketCount {
		result += 0.2
	} else if nameBetweenBacktickCount > 0 {
		result += 0.1
	}

	return result
}))
