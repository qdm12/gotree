package gotree

import "fmt"

type Node struct {
	value  string
	childs []*Node
}

func NewRoot(format string, args ...interface{}) *Node {
	return &Node{
		value: fmt.Sprintf(format, args...),
	}
}

func (n *Node) DeepCopy() *Node {
	nodeCopy := new(Node)
	nodeCopy.value = n.value
	if n.childs == nil {
		return nodeCopy
	}
	nodeCopy.childs = make([]*Node, len(n.childs))
	for i := range n.childs {
		nodeCopy.childs[i] = n.childs[i].DeepCopy()
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
	cp := node.DeepCopy()

	n.childs = append(n.childs, cp)
}

func (n *Node) ToLines() (lines []string) {
	const isRoot = true
	const isLast = false
	return toLines(n, isRoot, isLast)
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
