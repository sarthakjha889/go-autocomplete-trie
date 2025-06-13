package trie

import "fmt"

func Example() {
	t := New()
	t.Insert("Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")

	results := t.SearchAll("wdn")
	fmt.Println(results)

	results2 := t.SearchAll("tsd")
	fmt.Println(results2)

	// Output:
	// [Wednesday]
	// [Thursday Tuesday Wednesday]
}

func Example_noFeatures() {
	t := New().CaseSensitive().WithoutFuzzy().WithoutLevenshtein().WithoutNormalisation()
	t.Insert("Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")

	results := t.SearchAll("t")
	fmt.Println(results)

	results2 := t.SearchAll("T")
	fmt.Println(results2)

	// Output:
	// []
	// [Thursday Tuesday]
}

func Example_metadata() {
	t := NewG[struct{ ID int }]()
	t.Insert("iPhone", struct{ ID int }{1})
	for _, hit := range t.SearchAll("iphone") {
		fmt.Println(hit.Word, hit.Meta.ID)
	}
	// Output:
	// iPhone 1
}
