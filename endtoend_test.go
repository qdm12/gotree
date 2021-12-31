package gotree

import (
	"strings"
	"testing"
)

func Test_EndToEnd(t *testing.T) {
	t.Parallel()

	root := NewRoot("root")

	_ = root.Appendf("a")
	nodeB := root.Appendf("b")
	_ = nodeB.Appendf("x")

	nodeC := root.Appendf("c")
	nodeD := nodeC.Appendf("d")
	_ = nodeD.Appendf("p")
	_ = nodeD.Appendf("z")

	lines := root.ToLines()

	expectedLines := []string{
		"root",
		"├── a",
		"├── b",
		"|   └── x",
		"└── c",
		"    └── d",
		"        ├── p",
		"        └── z",
	}

	expected := strings.Join(expectedLines, "\n")
	actual := strings.Join(lines, "\n")

	if expected != actual {
		t.Errorf("actual result does not match expected result:\nActual:\n%s\nExpected:\n%s",
			actual, expected)
	}
}
