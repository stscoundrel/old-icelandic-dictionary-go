package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionariesHaveExpectedAmountOfEntries(t *testing.T) {
	expected_entries := 29951
	defaultResult, _ := GetDictionary()
	noMarkupResult, _ := GetNoMarkupDictionary()

	assert.Equal(t, expected_entries, len(defaultResult))
	assert.Equal(t, expected_entries, len(noMarkupResult))
	assert.Equal(t, len(defaultResult), len(noMarkupResult))
}

func TestDefaultDictionaryHasExpectedEntries(t *testing.T) {
	result, _ := GetDictionary()

	expected1 := DictionaryEntry{
		Headword: "afbindi",
		Definitions: []string{
			"n. <i>constipation</i>.",
		},
	}

	expected2 := DictionaryEntry{
		Headword: "andstreymr",
		Definitions: []string{
			"a. <i>strongly adverse</i> (andstreym ørlög); Sighvatr var heldr ~ um eptirmálin, <i>hard to come to terms with</i>.",
		},
	}

	expected3 := DictionaryEntry{
		Headword: "undanhald",
		Definitions: []string{
			"n. <i>flight</i>.",
		},
	}

	assert.Equal(t, expected1, result[14])
	assert.Equal(t, expected2, result[1000])
	assert.Equal(t, expected3, result[25000])
}

func TestNoMarkupDictionaryHasExpectedEntries(t *testing.T) {
	result, _ := GetNoMarkupDictionary()

	expected1 := DictionaryEntry{
		Headword: "afbindi",
		Definitions: []string{
			"n. constipation.",
		},
	}

	expected2 := DictionaryEntry{
		Headword: "andstreymr",
		Definitions: []string{
			"a. strongly adverse (andstreym ørlög); Sighvatr var heldr ~ um eptirmálin, hard to come to terms with.",
		},
	}

	expected3 := DictionaryEntry{
		Headword: "undanhald",
		Definitions: []string{
			"n. flight.",
		},
	}

	assert.Equal(t, expected1, result[14])
	assert.Equal(t, expected2, result[1000])
	assert.Equal(t, expected3, result[25000])
}
