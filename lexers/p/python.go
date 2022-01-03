package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Python lexer.
var Python = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Python",
		Aliases:   []string{"python", "py", "sage", "python3", "py3"},
		Filenames: []string{"*.py", "*.pyi", "*.pyw", "*.jy", "*.sage", "*.sc", "SConstruct", "SConscript", "*.bzl", "BUCK", "BUILD", "BUILD.bazel", "WORKSPACE", "*.tac"},
		MimeTypes: []string{"text/x-python", "application/x-python", "text/x-python3", "application/x-python3"},
	},
	embedded,
	"embedded/python.xml",
))
