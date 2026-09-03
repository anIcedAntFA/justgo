package exercises

import (
	"testing"

	"github.com/anIcedAntFA/justgo/content/part-1-foundation/11-packages-and-modules/exercises/catalog"
)

func TestCatalogVisibility(t *testing.T) {
	t.Skip("Chapter 11 exercise: implement package catalog, then delete this Skip")

	cases := []struct {
		name    string
		price   int
		wantSKU string
	}{
		{"Go Mug", 12, "go-mug"},
		{"Sticker Pack", 5, "sticker-pack"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := catalog.New(tc.name, tc.price)
			if p.Name != tc.name {
				t.Errorf("Name = %q, want %q", p.Name, tc.name)
			}
			if p.Price != tc.price {
				t.Errorf("Price = %d, want %d", p.Price, tc.price)
			}
			if got := p.SKU(); got != tc.wantSKU {
				t.Errorf("SKU() = %q, want %q", got, tc.wantSKU)
			}
		})
	}

	// The boundary in action — both lines below are COMPILE errors, because they
	// reach for identifiers package catalog does not export. Uncomment either to
	// watch the compiler enforce visibility from outside the package:
	//
	//	_ = catalog.Product{sku: "x"}   // ❌ sku is unexported
	//	_ = catalog.slugify("Go Mug")   // ❌ slugify is unexported
}
