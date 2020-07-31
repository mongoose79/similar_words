package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortString(t *testing.T) {
	res := SortString("dabc")
	assert.Equal(t, "abcd", res)

	res = SortString("")
	assert.Equal(t, "", res)
}

func TestIsStringInSlice(t *testing.T) {
	res := IsStringInSlice("abc", []string{"bac", "ab", "abc", "abcd"})
	assert.True(t, res)

	res = IsStringInSlice("abcf", []string{"bac", "ab", "abc", "abcd"})
	assert.False(t, res)

	res = IsStringInSlice("", []string{"bac", "ab", "abc", "abcd"})
	assert.False(t, res)

	res = IsStringInSlice("abc", []string{})
	assert.False(t, res)

	res = IsStringInSlice("abc", nil)
	assert.False(t, res)
}

func TestRemoveStrFromSlice(t *testing.T) {
	res := RemoveStrFromSlice([]string{"abc", "bac", "ab", "abcde"}, "bac")
	assert.Equal(t, []string{"abc", "ab", "abcde"}, res)

	res = RemoveStrFromSlice([]string{"abc", "bac", "ab", "abcde"}, "bacf")
	assert.Equal(t, []string{"abc", "bac", "ab", "abcde"}, res)
}
