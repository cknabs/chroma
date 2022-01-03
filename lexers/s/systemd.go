package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var SYSTEMD = internal.Register(MustNewXMLLexer(
	&Config{
		Name:    "SYSTEMD",
		Aliases: []string{"systemd"},
		// Suspects: man systemd.index | grep -E 'systemd\..*configuration'
		Filenames: []string{"*.automount", "*.device", "*.dnssd", "*.link", "*.mount", "*.netdev", "*.network", "*.path", "*.scope", "*.service", "*.slice", "*.socket", "*.swap", "*.target", "*.timer"},
		MimeTypes: []string{"text/plain"},
	},
	embedded,
	"embedded/systemd.xml",
))
