package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var ArmAsm = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "ArmAsm",
		Aliases:   []string{"armasm"},
		EnsureNL:  true,
		Filenames: []string{"*.s", "*.S"},
		MimeTypes: []string{"text/x-armasm", "text/x-asm"},
	},
	embedded,
	"embedded/armasm.xml",
))
