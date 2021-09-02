package tree

import (
	"path/filepath"
	"strings"
)

func uniqueCleanPaths(paths []string, rootPath string) (outPaths []string) {
	outPaths = make([]string, len(paths))
	for i := range paths {
		outPaths[i] = cleanPath(paths[i], rootPath)
	}

	outPaths = uniquePaths(outPaths)

	if len(outPaths) == 0 {
		return nil
	}
	return outPaths
}

func cleanPath(path, rootPath string) (cleanedPath string) {
	// Make path absolute with current path
	if !strings.HasPrefix(path, "/") {
		path = filepath.Join(rootPath, path)
	}

	path = filepath.Clean(path)

	return path
}

// func extractDirectoryPaths(filepaths []string) (directoryPaths []string) {
// 	uniqueDirPaths := make(map[string]struct{})
// 	for _, path := range filepaths {
// 		parts := strings.Split(path, "/")[1:]
// 		if len(parts) == 1 {
// 			continue // only one file
// 		}
// 		for i := 1; i < len(parts); i++ {
// 			directoryPath := "/" + strings.Join(parts[:i], "/")
// 			uniqueDirPaths[directoryPath] = struct{}{}
// 		}
// 	}

// 	if len(uniqueDirPaths) == 0 {
// 		return nil
// 	}

// 	directoryPaths = make([]string, 0, len(uniqueDirPaths))
// 	for path := range uniqueDirPaths {
// 		directoryPaths = append(directoryPaths, path)
// 	}

// 	sortPaths(directoryPaths)

// 	return directoryPaths
// }

func uniquePaths(paths []string) (uniquePaths []string) {
	set := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		set[path] = struct{}{}
	}

	uniquePaths = make([]string, 0, len(set))
	for path := range set {
		uniquePaths = append(uniquePaths, path)
	}

	return uniquePaths
}
