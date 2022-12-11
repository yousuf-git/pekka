package tree

import "strings"

type node struct {
	value   string
	handler bool
	nodes   []*node
}

type Tree struct {
	root *node
}

func New() Tree {
	return Tree{root: &node{
		value:   "",
		handler: false,
		nodes:   nil,
	}}
}

func DummyTree() Tree {
	return Tree{root: &node{
		value:   "",
		handler: true,
		nodes: []*node{
			{
				value:   "categories",
				handler: true,
				nodes: []*node{
					{
						value:   ":id",
						handler: false,
						nodes: []*node{
							{
								value:   "skus",
								handler: true,
							},
						},
					},
				},
			},
			{
				value:   "skus",
				handler: true,
			},
		},
	}}
}

func (t *Tree) Insert(pattern string) {
	// Remove leading '/'.
	if len(pattern) > 0 && pattern[0] == '/' {
		pattern = pattern[1:]
	}

	// Root handler.
	if pattern == "" {
		t.root.handler = true
		return
	}

	ss := strings.Split(pattern, "/")

	cur := t.root
	i, d := 0, 0
	for {
		if d > len(ss)-1 {
			break
		}

		if i > len(cur.nodes)-1 {
			// Create a new entry on no entry.
			cur.nodes = append(cur.nodes, &node{
				value:   ss[d],
				handler: false,
				nodes:   nil,
			})
			i = 0
			d++
			cur = cur.nodes[len(cur.nodes)-1] // Set to newly added one.
			continue
		}

		n := cur.nodes[i]
		if n.value == ss[d] {
			i = 0
			d++
			cur = n
			continue
		}

		i++
	}

	cur.handler = true
}

func (t *Tree) Has(pattern string) bool {
	// Remove leading '/'.
	if len(pattern) > 0 && pattern[0] == '/' {
		pattern = pattern[1:]
	}

	pattern = removeQueryAndHash(pattern)
	ss := strings.Split(pattern, "/")

	// Root handler.
	if pattern == "" {
		if t.root.handler == true {
			return true
		}
		return false
	}

	i, d := 0, 0
	cur := t.root
	for {
		if cur.nodes == nil {
			break
		}
		if i > len(cur.nodes)-1 {
			break
		}
		if d > len(ss)-1 {
			break
		}

		n := cur.nodes[i]
		if ss[d] == n.value || n.value[0] == ':' {
			d++
			i = 0
			cur = n
			continue
		}
		i++
	}

	if d == len(ss) && cur.handler == true {
		return true
	}
	return false
}

func removeQueryAndHash(s string) string {
	for i, c := range s {
		if c == '?' || c == '#' {
			return s[:i]
		}
	}
	return s
}
