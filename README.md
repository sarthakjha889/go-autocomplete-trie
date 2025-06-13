<h1 align="center">Go-Autocomplete-Trie</h1>

<div align="center">
  An <code>autocompl...</code> library for Go by Vivino.
</div>

<br />

[![GoDoc](https://godoc.org/github.com/your-org/fuzzytrie?status.svg)](https://godoc.org/github.com/your-org/fuzzytrie) [![Build Status](https://travis-ci.com/Vivino/go-autocomplete-trie.svg?branch=master)](https://travis-ci.com/Vivino/go-autocomplete-trie)

## What Is it

Go-Autocomplete-Trie is a simple, configurable autocompletion library for Go. Simply build a dictionary with a slice of strings, optionally configure, and then search.

## How to Use

Make a default Trie like so: 

```t := trie.New()``` 

The default Trie has *fuzzy* search enabled, string *normalisation* enabled, a default *levenshtein* scheme and is *case insensitive* by default.

Next, just add some strings to the dictionary.

```
t.Insert("Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")
```

Next, search.

```
t.SearchAll("wdn")

-> []string{"Wednesday"}
```

Levenshtein is enabled by default.

```
t.SearchAll("urs")

-> []string{"Thursday", "Tuesday"}
```

To turn off the features...

```
t.WithoutLevenshtein().WithoutNormalisation().WithoutFuzzy().CaseSensitive()
```

Now...

```
t.SearchAll("urs")

-> []string{}

t.SearchAll("Thu")

-> []string{"Thursday"}
```

### Using metadata

The trie can store arbitrary metadata with each entry.

```go
type Product struct{ ID int; Price float64 }

t := trie.NewG[Product]()
t.Insert("iPhone", Product{1, 999})

if p, ok := t.Find("iPhone"); ok {
        fmt.Println(p.ID, p.Price)
}

for _, hit := range t.SearchAll("iphne") {
        fmt.Printf("~ %s -> %+v\n", hit.Word, hit.Meta)
}
```
