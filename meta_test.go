package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Product struct {
	ID    int
	Price float64
}

func TestMetadataFunctions(t *testing.T) {
	t.Run("InsertWithMeta and FindMeta", func(t *testing.T) {
		tr := New()
		tr.InsertWithMeta("iPhone", "phone")
		meta, ok := tr.FindMeta("iPhone")
		assert.True(t, ok)
		assert.Equal(t, "phone", meta)
	})

	t.Run("BulkInsertWithMeta", func(t *testing.T) {
		tr := New()
		tr.BulkInsertWithMeta(map[string]interface{}{"ipad": 1, "mac": 2})
		m, ok := tr.FindMeta("ipad")
		assert.True(t, ok)
		assert.Equal(t, 1, m)
	})

	t.Run("SearchAllMeta", func(t *testing.T) {
		tr := New()
		tr.InsertWithMeta("iPhone", "device")
		hits := tr.SearchAllMeta("iphone")
		assert.Equal(t, 1, len(hits))
		assert.Equal(t, "iPhone", hits[0].Word)
		assert.Equal(t, "device", hits[0].Meta)
	})

	t.Run("Generic wrapper", func(t *testing.T) {
		g := NewG[Product]()
		g.Insert("iPhone", Product{ID: 1, Price: 999})
		p, ok := g.Find("iPhone")
		assert.True(t, ok)
		assert.Equal(t, 1, p.ID)
		res := g.SearchAll("iPhone")
		assert.Equal(t, 1, len(res))
		assert.Equal(t, 999.0, res[0].Meta.Price)
	})

	t.Run("Delete with metadata", func(t *testing.T) {
		tr := New()
		tr.InsertWithMeta("ipad", 1)
		tr.Delete("ipad")
		_, ok := tr.FindMeta("ipad")
		assert.False(t, ok)
		assert.Empty(t, tr.SearchAllMeta("ipad"))
	})

	t.Run("Fuzzy collisions", func(t *testing.T) {
		tr := New()
		tr.InsertWithMeta("iPhone", "A")
		tr.InsertWithMeta("iPhobe", "B")
		hits := tr.SearchAllMeta("iphoe")
		assert.Equal(t, 2, len(hits))
		m := map[string]interface{}{hits[0].Word: hits[0].Meta, hits[1].Word: hits[1].Meta}
		assert.Equal(t, "A", m["iPhone"])
		assert.Equal(t, "B", m["iPhobe"])
	})
}
