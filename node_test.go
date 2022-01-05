package gotree

import (
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		value string
		root  *Node
	}{
		"no value": {
			root: &Node{},
		},
		"with value": {
			value: "value",
			root: &Node{
				value: "value",
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			root := New(testCase.value)

			if !reflect.DeepEqual(testCase.root, root) {
				t.Errorf("actual result does not match expected result:\nActual:\n%#v\nExpected:\n%#v",
					root, testCase.root)
			}
		})
	}
}

func Test_Node_DeepCopy(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		original *Node
		copied   *Node
	}{
		"empty node": {
			original: &Node{},
			copied:   &Node{},
		},
		"node": {
			original: &Node{
				value: "value",
				childs: []*Node{
					{value: "child 1"},
					{value: "child 2"},
				},
			},
			copied: &Node{
				value: "value",
				childs: []*Node{
					{value: "child 1"},
					{value: "child 2"},
				},
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			copied := testCase.original.deepCopy()

			if !reflect.DeepEqual(testCase.copied, copied) {
				t.Errorf("actual result does not match expected result:\nActual:\n%#v\nExpected:\n%#v",
					copied, testCase.copied)
				return
			}

			copied.value += "x"
			for i := range copied.childs {
				copied.childs[i].value += "x"
			}

			if len(copied.childs) > 0 && reflect.DeepEqual(copied.childs, testCase.original.childs) {
				t.Error("modified copied childs modified original node childs")
			}
			if reflect.DeepEqual(copied.value, testCase.original.value) {
				t.Error("modified copied value modified original value")
			}
		})
	}
}

func Test_Node_Appendf(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		originalNode *Node
		format       string
		args         []interface{}
		newNode      *Node
		expectedNode *Node
	}{
		"empty node": {
			originalNode: &Node{},
			format:       "%s %d",
			args:         []interface{}{"hello", 1},
			newNode: &Node{
				value: "hello 1",
			},
			expectedNode: &Node{
				childs: []*Node{
					{
						value: "hello 1",
					},
				},
			},
		},
		"node with children": {
			originalNode: &Node{
				childs: []*Node{
					{
						value: "A",
					},
					{
						value: "B",
					},
				},
			},
			format: "%s %d",
			args:   []interface{}{"hello", 1},
			newNode: &Node{
				value: "hello 1",
			},
			expectedNode: &Node{
				childs: []*Node{
					{
						value: "A",
					},
					{
						value: "B",
					},
					{
						value: "hello 1",
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			newNode := testCase.originalNode.Appendf(testCase.format, testCase.args...)

			if !reflect.DeepEqual(testCase.newNode, newNode) {
				t.Errorf("actual new node does not match expected new node:\nActual:\n%#v\nExpected:\n%#v",
					newNode, testCase.newNode)
			}

			if !reflect.DeepEqual(testCase.expectedNode, testCase.originalNode) {
				t.Errorf("actual node does not match expected node:\nActual:\n%#v\nExpected:\n%#v",
					testCase.originalNode, testCase.expectedNode)
			}
		})
	}
}

func Test_Node_AppendNode(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		originalNode *Node
		appendedNode *Node
		expectedNode *Node
	}{
		"empty node": {
			originalNode: &Node{},
			appendedNode: &Node{
				value: "appended",
			},
			expectedNode: &Node{
				childs: []*Node{{
					value: "appended",
				}},
			},
		},
		"node with children": {
			originalNode: &Node{
				childs: []*Node{
					{
						value: "A",
					},
					{
						value: "B",
					},
				},
			},
			appendedNode: &Node{
				value: "appended",
			},
			expectedNode: &Node{
				childs: []*Node{
					{
						value: "A",
					},
					{
						value: "B",
					},
					{
						value: "appended",
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testCase.originalNode.AppendNode(testCase.appendedNode)

			if !reflect.DeepEqual(testCase.expectedNode, testCase.originalNode) {
				t.Errorf("actual new node does not match expected new node:\nActual:\n%#v\nExpected:\n%#v",
					testCase.originalNode, testCase.expectedNode)
			}
		})
	}
}

func Test_Node_String(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		node Node
		s    string
	}{
		"nil node": {},
		"single node": {
			node: Node{
				value: "value",
			},
			s: "value",
		},
		"node with children": {
			node: Node{
				value: "value",
				childs: []*Node{
					{value: "a", childs: []*Node{
						{value: "c"},
						{value: "d"},
					}},
					{value: "b"},
				},
			},
			s: `value
├── a
|   ├── c
|   └── d
└── b`,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := testCase.node.String()

			if testCase.s != s {
				t.Errorf("actual result does not match expected result:\nActual:\n%s\nExpected:\n%s",
					s, testCase.s)
			}
		})
	}

	t.Run("untyped nil node", func(t *testing.T) {
		t.Parallel()

		getNode := func() *Node {
			return nil
		}

		s := getNode().String()

		if s != "" {
			t.Errorf("string is not empty: %s", s)
		}
	})
}
