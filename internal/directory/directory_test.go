package directory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Parallel()

	const name = "name"

	intf := New(name)

	impl, ok := intf.(*directory)
	require.True(t, ok)

	expected := &directory{
		name: name,
	}
	assert.Equal(t, expected, impl)
}

func Test_directory_Name(t *testing.T) {
	t.Parallel()

	const expected = "name"
	d := &directory{name: expected}
	actual := d.Name()

	assert.Equal(t, expected, actual)
}

func Test_directory_Files(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		dir   *directory
		files []string
	}{
		"no files": {
			dir: &directory{},
		},
		"files in dir": {
			dir: &directory{
				files: []string{"a", "b"},
			},
			files: []string{"a", "b"},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			files := testCase.dir.Files()

			assert.Equal(t, testCase.files, files)
		})
	}

	t.Run("files returned do not modify directory", func(t *testing.T) {
		t.Parallel()

		dir := &directory{
			files: []string{"a"},
		}

		files := dir.Files()
		files[0] = "b"

		assert.Equal(t, []string{"a"}, dir.files)
	})
}

func Test_directory_Directories(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		dir         *directory
		directories []Directory
	}{
		"no directories": {
			dir: &directory{},
		},
		"directories in dir": {
			dir: &directory{
				directories: []Directory{
					&directory{name: "a"},
					&directory{name: "b"},
				},
			},
			directories: []Directory{
				&directory{name: "a"},
				&directory{name: "b"},
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			directories := testCase.dir.Directories()

			assert.Equal(t, testCase.directories, directories)
		})
	}

	t.Run("directories returned do not modify directory", func(t *testing.T) {
		t.Parallel()

		dir := &directory{
			directories: []Directory{&directory{name: "a"}},
		}

		directories := dir.Directories()
		directories[0] = nil

		assert.Equal(t, []Directory{&directory{name: "a"}}, dir.directories)
	})
}

func Test_directory_AddFile(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		initialDir  *directory
		file        string
		expectedDir *directory
	}{
		"first file": {
			initialDir: &directory{
				fileNames: map[string]struct{}{},
			},
			file: "a",
			expectedDir: &directory{
				fileNames: map[string]struct{}{"a": {}},
				files:     []string{"a"},
			},
		},
		"second file": {
			initialDir: &directory{
				fileNames: map[string]struct{}{"b": {}},
				files:     []string{"b"},
			},
			file: "a",
			expectedDir: &directory{
				fileNames: map[string]struct{}{"a": {}, "b": {}},
				files:     []string{"a", "b"},
			},
		},
		"duplicate file": {
			initialDir: &directory{
				fileNames: map[string]struct{}{"a": {}},
				files:     []string{"a"},
			},
			file: "a",
			expectedDir: &directory{
				fileNames: map[string]struct{}{"a": {}},
				files:     []string{"a"},
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			dir := testCase.initialDir

			dir.AddFile(testCase.file)

			assert.Equal(t, testCase.expectedDir, dir)
		})
	}
}

func Test_directory_AddDirectory(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		initialDir  *directory
		directory   *directory
		expectedDir *directory
	}{
		"first directory": {
			initialDir: &directory{
				directoryNames: map[string]struct{}{},
			},
			directory: &directory{name: "a"},
			expectedDir: &directory{
				directoryNames: map[string]struct{}{"a": {}},
				directories:    []Directory{&directory{name: "a"}},
			},
		},
		"second directory": {
			initialDir: &directory{
				directoryNames: map[string]struct{}{"b": {}},
				directories:    []Directory{&directory{name: "b"}},
			},
			directory: &directory{name: "a"},
			expectedDir: &directory{
				directoryNames: map[string]struct{}{"a": {}, "b": {}},
				directories:    []Directory{&directory{name: "a"}, &directory{name: "b"}},
			},
		},
		"duplicate directory": {
			initialDir: &directory{
				directoryNames: map[string]struct{}{"a": {}},
				directories:    []Directory{&directory{name: "a"}},
			},
			directory: &directory{name: "a"},
			expectedDir: &directory{
				directoryNames: map[string]struct{}{"a": {}},
				directories:    []Directory{&directory{name: "a"}},
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			dir := testCase.initialDir

			dir.AddDirectory(testCase.directory)

			assert.Equal(t, testCase.expectedDir, dir)
		})
	}
}

func Test_directory_Lines(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		dir   *directory
		lines []string
	}{
		"empty directory": {
			dir: &directory{},
		},
		"root directory": {
			dir: &directory{
				name: "/",
			},
			lines: []string{"/"},
		},
		"directory under root": {
			dir: &directory{
				name: "/",
				directories: []Directory{
					&directory{name: "a"},
				},
			},
			lines: []string{
				"/",
				"└── a/",
			},
		},
		"files under root": {
			dir: &directory{
				name:  "/",
				files: []string{"a", "b"},
			},
			lines: []string{
				"/",
				"├── a",
				"└── b",
			},
		},
		"nested directories": {
			dir: &directory{
				name: "/",
				directories: []Directory{
					&directory{
						name:        "a",
						directories: []Directory{&directory{name: "b"}},
					},
				},
			},
			lines: []string{
				"/",
				"└── a/",
				"    └── b/",
			},
		},
		"directories and files": {
			dir: &directory{
				name: "/",
				directories: []Directory{
					&directory{
						name: "a1",
						directories: []Directory{
							&directory{
								name:  "a2",
								files: []string{"f3"},
							},
							&directory{
								name:  "a3",
								files: []string{"f4"},
							},
						},
					},
					&directory{
						name:        "b1",
						directories: []Directory{&directory{name: "b2"}},
					},
				},
				files: []string{"f1", "f2"},
			},
			lines: []string{
				"/",
				"├── f1",
				"├── f2",
				"├── a1/",
				"|   ├── a2/",
				"|   |   └── f3",
				"|   └── a3/",
				"|       └── f4",
				"└── b1/",
				"    └── b2/",
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			lines := testCase.dir.Lines()

			assert.Equal(t, testCase.lines, lines)
		})
	}
}
