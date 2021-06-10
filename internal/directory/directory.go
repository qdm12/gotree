package directory

import "sort"

type Directory interface {
	Name() string
	Files() (names []string)
	Directories() (directories []Directory)
	AddDirectory(directory Directory)
}

type directory struct {
	name        string
	files       []string
	directories []Directory
}

func New(name string, files []string) Directory {
	return &directory{
		name:  name,
		files: files,
	}
}

func (d *directory) Name() string { return d.name }

func (d *directory) Files() (names []string) { return d.files }

func (d *directory) Directories() (directories []Directory) { return d.directories }

func (d *directory) AddDirectory(directory Directory) {
	d.directories = append(d.directories, directory)

	// sort the directories alphabetically
	sort.Slice(d.directories, func(i, j int) bool {
		return d.directories[i].Name() < d.directories[j].Name()
	})
}
