package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tree_ToLines(t *testing.T) {
	t.Parallel()

	// More cases:
	// - extrapolate directories from file path

	// Assume all filePaths and directoryPaths have been through sanitizePaths
	// so they are all already alphabetically sorted.
	testCases := map[string]struct {
		tree  Tree
		lines []string
	}{
		"empty tree": {
			tree:  &tree{},
			lines: []string{"."},
		},
		"absolute file path": {
			tree: &tree{
				filePaths: []string{"/file"},
			},
			lines: []string{
				"/",
				"└── file",
			},
		},
		"relative file path": {
			tree: &tree{
				filePaths: []string{"file"},
			},
			lines: []string{
				".",
				"└── file",
			},
		},
		"absolute directory path": {
			tree: &tree{
				directoryPaths: []string{"/directory"},
			},
			lines: []string{
				"/",
				"└── directory",
			},
		},
		"relative directory path": {
			tree: &tree{
				directoryPaths: []string{"directory"},
			},
			lines: []string{
				".",
				"└── directory/",
			},
		},
		"files before directories": {
			tree: &tree{
				filePaths:      []string{"/b"},
				directoryPaths: []string{"/a"},
			},
			lines: []string{
				"/",
				"├── b",
				"├── a/",
			},
		},
		"mix of things": {
			tree: &tree{
				filePaths: []string{
					"/dir/dirA/file",
					"/dir/dirB/dir/file",
					"/file",
				},
				directoryPaths: []string{
					"/dir/",
					"/dirA/",
					"/dirB/",
					"/dirB/dir",
				},
			},
			lines: []string{
				"/",
				"├── file",
				"└── dir/",
				"    ├── dirA/",
				"    |   └── file",
				"    └── dirB/",
				"        └── dir/",
				"            └── file",
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			tree := testCase.tree
			lines := tree.ToLines()

			assert.Equal(t, testCase.lines, lines)
		})
	}
}
