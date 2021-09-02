package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tree(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		rootPath       string
		directoryPaths []string
		filePaths      []string
		lines          []string
	}{
		"no path": {},
		"relative path with root": {
			rootPath:  "/",
			filePaths: []string{"f1"},
			lines: []string{
				"/",
				"└── f1",
			},
		},
		"relative path with custom root path": {
			rootPath:  "/test",
			filePaths: []string{"f1"},
			lines: []string{
				".",
				"└── f1",
			},
		},
		// "absolute file path": {
		// 	filePaths: []string{"/file"},
		// 	lines: []string{
		// 		"/",
		// 		"└── file",
		// 	},
		// },
		// "relative file path": {
		// 	filePaths: []string{"file"},
		// 	lines: []string{
		// 		".",
		// 		"└── file",
		// 	},
		// },
		// "absolute directory path": {
		// 	directoryPaths: []string{"/directory"},
		// 	lines: []string{
		// 		"/",
		// 		"└── directory/",
		// 	},
		// },
		// "relative directory path": {
		// 	directoryPaths: []string{"directory"},
		// 	lines: []string{
		// 		".",
		// 		"└── directory/",
		// 	},
		// },
		// "files before directories": {
		// 	filePaths:      []string{"/b"},
		// 	directoryPaths: []string{"/a"},
		// 	lines: []string{
		// 		"/",
		// 		"├── b",
		// 		"├── a/",
		// 	},
		// },
		// "mix of things": {
		// 	filePaths: []string{
		// 		"/dir/dirA/file",
		// 		"/dir/dirB/dir/file",
		// 		"/file",
		// 	},
		// 	directoryPaths: []string{
		// 		"/dirC/",
		// 		"/dirA/dirD",
		// 	},
		// 	lines: []string{
		// 		"/",
		// 		"├── file",
		// 		"└── dir/",
		// 		"    ├── dirA/",
		// 		"    |   └── file",
		// 		"    └── dirB/",
		// 		"        └── dir/",
		// 		"            └── file",
		// 	},
		// },
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			tree := New(testCase.rootPath, testCase.directoryPaths, testCase.filePaths)
			lines := tree.Lines()

			assert.Equal(t, testCase.lines, lines)
		})
	}
}
