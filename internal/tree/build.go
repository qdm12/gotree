package tree

import (
	"strings"

	"github.com/qdm12/gotree/internal/directory"
)

func buildTree(rootPath string, directoryPaths, filePaths []string) (
	root directory.Directory) {
	root = directory.New(rootPath)

	// Set directories only
	for _, path := range directoryPaths {
		path = strings.TrimPrefix(path, rootPath)

		if path == "" {
			continue
		}

		parts := strings.Split(path, "/")[1:]

		currentDir := root
		for _, part := range parts {
			newDir := directory.New(part)
			currentDir.AddDirectory(newDir)
			currentDir = newDir
		}
	}

	// Set directories and files
	for _, path := range filePaths {
		path = strings.TrimPrefix(path, rootPath)

		parts := strings.Split(path, "/")[1:]
		if len(parts) == 0 { // no slash
			continue
		}

		currentDir := root
		for i := 0; i < len(parts)-1; i++ {
			directoryName := parts[i]
			existingDir, ok := currentDir.GetDirectory(directoryName)
			if ok {
				currentDir = existingDir
				continue
			}
			newDir := directory.New(parts[i])
			currentDir.AddDirectory(newDir)
			currentDir = newDir
		}

		fileName := parts[len(parts)-1]
		currentDir.AddFile(fileName)
	}

	return root
}
