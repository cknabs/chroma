package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cmake lexer.
var Cmake = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "CMake",
		Aliases:   []string{"cmake"},
		Filenames: []string{"*.cmake", "CMakeLists.txt"},
		MimeTypes: []string{"text/x-cmake"},
	},
	embedded,
	"embedded/cmake.xml",
))
