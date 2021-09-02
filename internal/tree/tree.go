package tree

import (
	"path/filepath"

	"github.com/qdm12/gotree/internal/directory"
)

type Tree interface {
	Lines() (lines []string)
}

type tree struct {
	root directory.Directory
}

func New(rootPath string, directoryPaths, filePaths []string) Tree {
	rootPath, err := filepath.Abs(rootPath)
	if err != nil {
		panic(err)
	}
	directoryPaths = uniqueCleanPaths(directoryPaths, rootPath)
	filePaths = uniqueCleanPaths(filePaths, rootPath)
	root := buildTree(rootPath, directoryPaths, filePaths)
	return &tree{
		root: root,
	}
}
