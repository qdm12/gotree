package directory

import (
	"sort"
	"strings"
)

type Directory interface {
	Name() string
	Files() (names []string)
	Directories() (directories []Directory)
	AddFile(file string)
	AddDirectory(directory Directory)
	GetDirectory(name string) (dir Directory, ok bool)
	Lines() (lines []string)
	recursiveLines() (lines []string)
}

type directory struct {
	name        string
	files       []string
	directories []Directory
	// for deduplication
	fileNames      map[string]struct{}
	directoryNames map[string]struct{}
}

func New(name string) Directory {
	return &directory{
		name:           name,
		fileNames:      make(map[string]struct{}),
		directoryNames: make(map[string]struct{}),
	}
}

func (d *directory) Name() string { return d.name }

func (d *directory) Files() (files []string) {
	if len(d.files) == 0 {
		return nil
	}
	files = make([]string, len(d.files))
	copy(files, d.files)
	return files
}

func (d *directory) Directories() (directories []Directory) {
	if len(d.directories) == 0 {
		return nil
	}
	directories = make([]Directory, len(d.directories))
	copy(directories, d.directories)
	return directories
}

func (d *directory) AddFile(file string) {
	if _, ok := d.fileNames[file]; ok {
		return // already exists
	}
	d.fileNames[file] = struct{}{}

	i := sort.SearchStrings(d.files, file)
	d.files = append(d.files, "")
	copy(d.files[i+1:], d.files[i:])
	d.files[i] = file
}

func (d *directory) AddDirectory(directory Directory) {
	name := directory.Name()

	if _, ok := d.directoryNames[name]; ok {
		return // already exists
	}
	d.directoryNames[name] = struct{}{}

	i := sort.Search(len(d.directories), func(i int) bool {
		return d.directories[i].Name() >= name
	})
	d.directories = append(d.directories, nil)
	copy(d.directories[i+1:], d.directories[i:])
	d.directories[i] = directory
}

func (d *directory) GetDirectory(name string) (dir Directory, ok bool) {
	if _, ok := d.directoryNames[name]; !ok {
		return nil, false
	}

	i := sort.Search(len(d.directories), func(i int) bool {
		return d.directories[i].Name() == name
	})

	return d.directories[i], true
}

func (d *directory) Lines() (lines []string) {
	if d.name == "/" {
		lines = append(lines, "/")
	} else if strings.HasPrefix(d.name, "/") {
		lines = append(lines, ".") // current directory
	}
	lines = append(lines, d.recursiveLines()...)
	return lines
}

func (d *directory) recursiveLines() (lines []string) {
	const prefix = "├── "
	const lastPrefix = "└── "

	i := 0
	total := len(d.files) + len(d.directories)

	// Files first
	for _, name := range d.files {
		prefix := prefix
		if i == total-1 {
			prefix = lastPrefix
		}
		lines = append(lines, prefix+name)
		i++
	}

	// Directories
	const dirSuffix = "/"
	for _, directory := range d.directories {
		last := i == total-1
		prefix := prefix
		if last {
			prefix = lastPrefix
		}
		line := prefix + directory.Name() + dirSuffix
		lines = append(lines, line)

		const prefixLength = 4 // cannot use len(prefix)
		linePrefix := "|" + strings.Repeat(" ", prefixLength-1)
		if last {
			linePrefix = strings.Repeat(" ", prefixLength)
		}
		for _, line := range directory.recursiveLines() {
			lines = append(lines, linePrefix+line)
		}

		i++
	}

	return lines
}
