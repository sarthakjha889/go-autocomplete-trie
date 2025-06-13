package main

import (
	"fmt"
	trie "github.com/sarthakjha889/go-autocomplete-trie"
)

type Product struct {
	ID    int
	Price float64
}

func main() {
	t := trie.NewG[Product]()
	t.Insert("iPhone", Product{ID: 1, Price: 999})
	t.Insert("iPad", Product{ID: 2, Price: 799})

	if p, ok := t.Find("iPhone"); ok {
		fmt.Println("Exact:", p.ID, p.Price)
	}

	for _, hit := range t.SearchAll("iphne") {
		fmt.Printf("~ %s → %+v\n", hit.Word, hit.Meta)
	}
}
