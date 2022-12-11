package main

import (
	"fmt"
	"pekka/tree"
)

func main() {
	t := tree.New()
	t.Insert("/")
	t.Insert("/categories/name")
	t.Insert("/categories/:categoryId/skus/:skuId")
	t.Insert("/categories/:categoryId/tags/analytics/monthly")
	t.Insert("/skus/name")
	test(t, []string{
		"",
		"/categories",
		"/categories/name",
		"/categories/100/skus",
		"/categories/200/tags/analytics/monthly",
		"/skus/name",
	})
}

func test(t tree.Tree, patterns []string) {
	for _, pattern := range patterns {
		b := t.Has(pattern)
		fmt.Println(pattern, b)
	}
}
