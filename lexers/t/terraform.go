package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Terraform lexer.
var Terraform = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Terraform",
		Aliases:   []string{"terraform", "tf"},
		Filenames: []string{"*.tf"},
		MimeTypes: []string{"application/x-tf", "application/x-terraform"},
	},
	embedded,
	"embedded/terraform.xml",
))
