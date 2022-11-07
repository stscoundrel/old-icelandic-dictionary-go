# Old Icelandic Dictionary

Old Icelandic dictionary for Golang. From "A Concise Dictionary of Old Icelandic" by Geir Zoëga

The dictionary consists of 29 000+ Old Icelandic words with English translations.

### Install

`go get github.com/stscoundrel/old-icelandic-dictionary-go`

### Usage

The dictionary comes in two variants:
- Default dictionary has html markup `<i>` and `<b>` to match look of the original book.
- No-markup version has the same content without any additional formatting tags.

```go
package main

import (
    "fmt"

    dictionary "github.com/stscoundrel/old-icelandic-dictionary-go"
)

func main() {
  defaultDictionary, defaltErr := dictionary.GetDictionary()
  noMarkupDictionary, noMarkupErr := dictionary.GetNoMarkupDictionary()
  
  // Error handling as you please.
  if defaltErr != nil {
    fmt.Println(err)
  }
  
  // Contains 29 000+ DictionaryEntries.
  for _, entry := range defaultDictionary {
    fmt.Println(entry.Headword)
  }
  
  // Headwords wont differ in dictionaries.
  fmt.Println(defaultDictionary[14].Headword)  // afbindi
  fmt.Println(noMarkupDictionary[14].Headword) // afbindi
  
  // But definitions markup will differ
  fmt.Println(defaultDictionary[14].Definitions[0])  // n. <i>constipation</i>.
  fmt.Println(noMarkupDictionary[14].Definitions[0]) // n. constipation.
}
```

The entries are structs of:

```go
type DictionaryEntry struct {
  Headword          string
  Definitions       []string
}

```

### About "A Concise Dictionary of Old Icelandic"

"A Concise Dictionary of Old Icelandic" dictionary was published in 1910 by Geir Zoëga, which leads to there being many public domain versions of the book available. Zoëgas attempt was to made easier-to-approach version of the more full Cleasby - Vigfusson dictionary, specifically for beginners and those interested in Old Icelandic prose writing.
