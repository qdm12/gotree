package tree

func (t *tree) Lines() (lines []string) {
	return t.root.Lines()
}
