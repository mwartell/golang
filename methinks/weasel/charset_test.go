package weasel

import (
	"testing"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected Charset
	}{
		{"abc", Charset{'a': true, 'b': true, 'c': true}},
		{"aabbcc", Charset{'a': true, 'b': true, 'c': true}},
		{"", Charset{}},
		{"123", Charset{'1': true, '2': true, '3': true}},
	}

	for _, test := range tests {
		result := FromString(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("FromString(%q) = %v; want %v", test.input, result, test.expected)
			continue
		}
		for k := range test.expected {
			if !result[k] {
				t.Errorf("FromString(%q) = %v; want %v", test.input, result, test.expected)
				break
			}
		}
	}
}

func TestCharsetString(t *testing.T) {
	tests := []struct {
		input    Charset
		expected string
	}{
		{Charset{'a': true, 'b': true, 'c': true}, "abc"},
		{Charset{'c': true, 'b': true, 'a': true}, "abc"},
		{Charset{}, ""},
		{Charset{'3': true, '1': true, '2': true}, "123"},
	}

	for _, test := range tests {
		result := test.input.String()
		if result != test.expected {
			t.Errorf("Charset.String() = %q; want %q", result, test.expected)
		}
	}
}
