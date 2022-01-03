package q

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Qml lexer.
var Qml = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "QML",
		Aliases:   []string{"qml", "qbs"},
		Filenames: []string{"*.qml", "*.qbs"},
		MimeTypes: []string{"application/x-qml", "application/x-qt.qbs+qml"},
		DotAll:    true,
	},
	embedded,
	"embedded/qml.xml",
))
