package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParserFindTitle(t *testing.T) {
	p := NewParser()

	for input, expectedTitle := range map[string]string{
		`
# Very healthy recipe
`: "Very healthy recipe",
		`
Pizza
=====
`: "Pizza",
		`
## Burger

`: "",
	} {
		meta, err := p.ParseRecipe([]byte(input))
		assert.NoError(t, err)
		require.NotNil(t, meta)
		assert.Equal(t, expectedTitle, meta.Title)
	}
}
