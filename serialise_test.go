package chroma

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	actual := MustNewLazyLexer(&Config{
		Name:      "PkgConfig",
		Aliases:   []string{"pkgconfig"},
		Filenames: []string{"*.pc"},
	}, func() Rules {
		return Rules{
			"root": {
				{`#.*$`, CommentSingle, nil},
				{`^(\w+)(=)`, ByGroups(NameAttribute, Operator), nil},
				{`^([\w.]+)(:)`, ByGroups(NameTag, Punctuation), Push("spvalue")},
				Include("interp"),
				{`[^${}#=:\n.]+`, Text, nil},
				{`.`, Text, nil},
			},
			"interp": {
				{`\$\$`, Text, nil},
				{`\$\{`, LiteralStringInterpol, Push("curly")},
			},
			"curly": {
				{`\}`, LiteralStringInterpol, Pop(1)},
				{`\w+`, NameAttribute, nil},
			},
			"spvalue": {
				Include("interp"),
				{`#.*$`, CommentSingle, Pop(1)},
				{`\n`, Text, Pop(1)},
				{`[^${}#\n]+`, Text, nil},
				{`.`, Text, nil},
			},
		}
	})
	data, err := Marshal(actual)
	require.NoError(t, err)
	expected, err := Unmarshal(data)
	require.NoError(t, err)
	require.Equal(t, expected.Config(), actual.Config())
	require.Equal(t, mustRules(t, expected), mustRules(t, actual))
}

func mustRules(t testing.TB, r *RegexLexer) Rules {
	t.Helper()
	rules, err := r.Rules()
	require.NoError(t, err)
	return rules
}

func TestRuleSerialisation(t *testing.T) {
	tests := []Rule{
		Include("String"),
		{`\d+`, Text, nil},
		{`"`, String, Push("String")},
	}
	for _, test := range tests {
		data, err := xml.Marshal(test)
		require.NoError(t, err)
		t.Log(string(data))
		actual := Rule{}
		err = xml.Unmarshal(data, &actual)
		require.NoError(t, err)
		require.Equal(t, test, actual)
	}
}

func TestRulesSerialisation(t *testing.T) {
	expected := Rules{
		"root": {
			{`#.*$`, CommentSingle, nil},
			{`^(\w+)(=)`, ByGroups(NameAttribute, Operator), nil},
			{`^([\w.]+)(:)`, ByGroups(NameTag, Punctuation), Push("spvalue")},
			Include("interp"),
			{`[^${}#=:\n.]+`, Text, nil},
			{`.`, Text, nil},
		},
		"interp": {
			{`\$\$`, Text, nil},
			{`\$\{`, LiteralStringInterpol, Push("curly")},
		},
		"curly": {
			{`\}`, LiteralStringInterpol, Pop(1)},
			{`\w+`, NameAttribute, nil},
		},
		"spvalue": {
			Include("interp"),
			{`#.*$`, CommentSingle, Pop(1)},
			{`\n`, Text, Pop(1)},
			{`[^${}#\n]+`, Text, nil},
			{`.`, Text, nil},
		},
	}
	data, err := xml.MarshalIndent(expected, "  ", "  ")
	require.NoError(t, err)
	re := regexp.MustCompile(`></[a-zA-Z]+>`)
	data = re.ReplaceAll(data, []byte(`/>`))
	fmt.Println(string(data))
	b := &bytes.Buffer{}
	w := gzip.NewWriter(b)
	fmt.Fprintln(w, string(data))
	w.Close()
	fmt.Println(len(data), b.Len())
	actual := Rules{}
	err = xml.Unmarshal(data, &actual)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}
