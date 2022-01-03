package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CPP lexer.
var CPP = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "C++",
		Aliases:   []string{"cpp", "c++"},
		Filenames: []string{"*.cpp", "*.hpp", "*.c++", "*.h++", "*.cc", "*.hh", "*.cxx", "*.hxx", "*.C", "*.H", "*.cp", "*.CPP"},
		MimeTypes: []string{"text/x-c++hdr", "text/x-c++src"},
		EnsureNL:  true,
	},
	embedded,
	"embedded/c++.xml",
))
