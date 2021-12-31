package gotree

import (
	"testing"
)

func Test_EndToEnd(t *testing.T) {
	t.Parallel()

	root := New("root")

	_ = root.Appendf("a")
	nodeB := root.Appendf("b")
	_ = nodeB.Appendf("x")

	nodeC := root.Appendf("c")
	nodeD := nodeC.Appendf("d")
	_ = nodeD.Appendf("p")
	_ = nodeD.Appendf("z")

	const expected = `root
├── a
├── b
|   └── x
└── c
    └── d
        ├── p
        └── z`
	s := root.String() //nolint:ifshort

	if expected != s {
		t.Errorf("actual result does not match expected result:\nActual:\n%s\nExpected:\n%s",
			s, expected)
	}
}
