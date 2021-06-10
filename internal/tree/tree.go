package tree

import (
	"sort"
	"strings"
)

type Tree interface {
	ToLines() (lines []string)
}

type tree struct {
	directoryPaths []string
	filePaths      []string
}

func New(directoryPaths, filePaths []string) Tree {
	return &tree{
		directoryPaths: sanitizePaths(directoryPaths),
		filePaths:      sanitizePaths(filePaths),
	}
}

func sanitizePaths(paths []string) []string {
	// sort them alphabetically
	sort.Slice(paths, func(i, j int) bool {
		return paths[i] < paths[j]
	})

	// remove the trailing slash
	for i := range paths {
		paths[i] = strings.TrimSuffix(paths[i], "/")
	}

	return paths
}
