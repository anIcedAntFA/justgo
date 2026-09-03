// Package catalog is the library under test for Exercise 1. It demonstrates Go's
// visibility rule: capitalization decides what code outside this package can use.
//
// Implement the exported API below. The test lives in the parent `exercises`
// package and imports this one, so it can reach ONLY the exported identifiers —
// that boundary is the whole point of the exercise.
//
// Note: this stub ships with no imports. You'll add the one `slugify` needs
// yourself — writing the import line is part of a Packages chapter.
package catalog

// Product is exported so other packages can name it. Two fields are exported
// (Name, Price); the third (sku) is unexported — private to package catalog.
type Product struct {
	Name  string
	Price int
	sku   string
}

// New builds a Product. It is named New, not NewProduct, to avoid stutter: callers
// write catalog.New, which already reads cleanly.
//
// TODO: return a *Product with the given name and price, and set the unexported
// sku field to slugify(name).
func New(name string, price int) *Product {
	return nil // TODO
}

// SKU is the exported reader for the unexported sku field — the only way code
// outside package catalog can see it.
//
// TODO: return the sku field.
func (p *Product) SKU() string {
	return "" // TODO
}

// slugify is an unexported helper — code outside package catalog cannot call it.
// Lowercase the name and replace spaces with dashes: "Go Mug" -> "go-mug".
//
// TODO: implement (import "strings"; strings.ToLower + strings.ReplaceAll is enough).
func slugify(name string) string {
	return "" // TODO
}
