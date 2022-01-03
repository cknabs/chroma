package b

import (
	"regexp"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TODO(moorereason): can this be factored away?
var bashAnalyserRe = regexp.MustCompile(`(?m)^#!.*/bin/(?:env |)(?:bash|zsh|sh|ksh)`)

// Bash lexer.
var Bash = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Bash",
		Aliases:   []string{"bash", "sh", "ksh", "zsh", "shell"},
		Filenames: []string{"*.sh", "*.ksh", "*.bash", "*.ebuild", "*.eclass", ".env", "*.env", "*.exheres-0", "*.exlib", "*.zsh", "*.zshrc", ".bashrc", "bashrc", ".bash_*", "bash_*", "zshrc", ".zshrc", "PKGBUILD"},
		MimeTypes: []string{"application/x-sh", "application/x-shellscript"},
	},
	embedded,
	"embedded/bash.xml",
).SetAnalyser(func(text string) float32 {
	if bashAnalyserRe.FindString(text) != "" {
		return 1.0
	}
	return 0.0
}))
