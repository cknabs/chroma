package x

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// XML lexer.
var XML = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "XML",
		Aliases:   []string{"xml"},
		Filenames: []string{"*.xml", "*.xsl", "*.rss", "*.xslt", "*.xsd", "*.wsdl", "*.wsf", "*.svg", "*.csproj", "*.vcxproj", "*.fsproj"},
		MimeTypes: []string{"text/xml", "application/xml", "image/svg+xml", "application/rss+xml", "application/atom+xml"},
		DotAll:    true,
	},
	embedded,
	"embedded/xml.xml",
))
