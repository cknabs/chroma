package y

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var YAML = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "YAML",
		Aliases:   []string{"yaml"},
		Filenames: []string{"*.yaml", "*.yml"},
		MimeTypes: []string{"text/x-yaml"},
	},
	embedded,
	"embedded/yaml.xml",
))
