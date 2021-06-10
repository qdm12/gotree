package directory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Parallel()

	const name = "name"
	files := []string{"a", "b"}

	intf := New(name, files)

	impl, ok := intf.(*directory)
	require.True(t, ok)

	expected := &directory{
		name:  name,
		files: files,
	}
	assert.Equal(t, expected, impl)
}

func Test_Name(t *testing.T) {
	t.Parallel()

	const expected = "name"
	d := &directory{name: expected}
	actual := d.Name()

	assert.Equal(t, expected, actual)
}
