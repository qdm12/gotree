package gotree

import (
	"fmt"
	"strings"
)

type Node struct {
	value  string
	childs []*Node
}

func New(value string) *Node {
	return &Node{
		value: value,
	}
}

func (n *Node) deepCopy() *Node {
	nodeCopy := new(Node)
	nodeCopy.value = n.value
	if n.childs == nil {
		return nodeCopy
	}
	nodeCopy.childs = make([]*Node, len(n.childs))
	for i := range n.childs {
		nodeCopy.childs[i] = n.childs[i].deepCopy()
	}
	return nodeCopy
}

func (n *Node) Appendf(format string, args ...interface{}) (newNode *Node) {
	newNode = &Node{
		value: fmt.Sprintf(format, args...),
	}

	n.childs = append(n.childs, newNode)

	return newNode
}

func (n *Node) AppendNode(node *Node) {
	if node == nil {
		// do not append nil node
		return
	}

	cp := node.deepCopy()

	n.childs = append(n.childs, cp)
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}

	const isRoot = true
	const isLast = false
	lines := toLines(n, isRoot, isLast)
	return strings.Join(lines, "\n")
}

func toLines(node *Node, isRoot, isLast bool) (lines []string) {
	var valuePrefix, childStart string
	switch {
	case isRoot:
	case isLast:
		childStart = "    "
		valuePrefix = "└── "
	default:
		childStart = "|   "
		valuePrefix = "├── "
	}
	lines = append(lines, valuePrefix+node.value)

	for i, child := range node.childs {
		isLast := i == len(node.childs)-1
		const isRoot = false
		for _, childLine := range toLines(child, isRoot, isLast) {
			lines = append(lines, childStart+childLine)
		}
	}

	return lines
}
