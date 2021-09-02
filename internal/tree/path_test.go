package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_cleanAndSortedPaths(t *testing.T) {
// 	t.Parallel()

// 	testCases := map[string]struct {
// 		paths         []string
// 		currentPath   string
// 		expectedPaths []string
// 	}{
// 		"no path": {
// 			currentPath: "/current",
// 		},
// 		"single path": {
// 			paths:         []string{"dir/file"},
// 			currentPath:   "/current",
// 			expectedPaths: []string{"/current/dir/file"},
// 		},
// 		"multiple paths": {
// 			paths:         []string{"/dir/file", "dir/file"},
// 			currentPath:   "/current",
// 			expectedPaths: []string{"/current/dir/file", "/dir/file"},
// 		},
// 	}

// 	for name, testCase := range testCases {
// 		testCase := testCase
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()

// 			paths := cleanAndSortedPaths(testCase.paths, testCase.currentPath)

// 			assert.Equal(t, testCase.expectedPaths, paths)
// 		})
// 	}
// }

func Test_cleanPath(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		path        string
		currentPath string
		cleanedPath string
	}{
		"empty": {
			currentPath: "/current",
			cleanedPath: "/current",
		},
		"dot": {
			path:        ".",
			currentPath: "/current",
			cleanedPath: "/current",
		},
		"root": {
			path:        "/",
			currentPath: "/current",
			cleanedPath: "/",
		},
		"trim trailing slash": {
			path:        "/dir/",
			cleanedPath: "/dir",
		},
		"relative path": {
			path:        "dir/file",
			currentPath: "/current",
			cleanedPath: "/current/dir/file",
		},
		"relative path with dot": {
			path:        "./dir/file",
			currentPath: "/current",
			cleanedPath: "/current/dir/file",
		},
		"absolute deep path": {
			path:        "/dir/file",
			currentPath: "/current",
			cleanedPath: "/dir/file",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cleanedPath := cleanPath(testCase.path, testCase.currentPath)

			assert.Equal(t, testCase.cleanedPath, cleanedPath)
		})
	}
}

// func Test_extractDirectoryPaths(t *testing.T) {
// 	t.Parallel()

// 	testCases := map[string]struct {
// 		filepaths      []string
// 		directoryPaths []string
// 	}{
// 		"no filepath": {},
// 		"top level filepath": {
// 			filepaths: []string{"/file"},
// 		},
// 		"single depth filepath": {
// 			filepaths:      []string{"/dir/file"},
// 			directoryPaths: []string{"/dir"},
// 		},
// 		"double depth filepath": {
// 			filepaths:      []string{"/dir1/dir2/file"},
// 			directoryPaths: []string{"/dir1", "/dir1/dir2"},
// 		},
// 	}

// 	for name, testCase := range testCases {
// 		testCase := testCase
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()

// 			directoryPaths := extractDirectoryPaths(testCase.filepaths)

// 			assert.Equal(t, testCase.directoryPaths, directoryPaths)
// 		})
// 	}
// }
