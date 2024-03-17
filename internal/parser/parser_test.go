package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParserFindTitle(t *testing.T) {
	p := NewParser()

	for _, testcase := range []struct {
		input         string
		expectedTitle string
		expectedError error
	}{
		{
			input: `
# Very healthy recipe
`,
			expectedTitle: "Very healthy recipe",
		},
		{
			input: `
Pizza
=====
`,
			expectedTitle: "Pizza",
		},
		{
			input: `
## Burger

Add some spice
`,
			expectedError: ErrMissingTitle,
		},
	} {
		meta, err := p.ParseRecipe([]byte(testcase.input))

		if testcase.expectedError != nil {
			assert.Nil(t, meta)
			assert.ErrorIs(t, err, testcase.expectedError)
		}

		if testcase.expectedTitle != "" {
			assert.NoError(t, err)
			require.NotNil(t, meta)
			assert.Equal(t, testcase.expectedTitle, meta.Title)
		}
	}
}

func TestParserWithValidFrontmatter(t *testing.T) {
	const input = `---
servings: 3
tags:
  - meat
source: https://example.org
...

# Recipe`

	expectedMeta := &RecipeMetadata{
		RecipeFrontmatter: &RecipeFrontmatter{
			Servings: 3,
			Tags:     []string{"meat"},
			Source:   "https://example.org",
		},
		Title: "Recipe",
	}

	p := NewParser()

	meta, err := p.ParseRecipe([]byte(input))
	assert.NoError(t, err)
	assert.EqualValues(t, expectedMeta, meta)
}

func TestParserWithInvalidFrontmatter(t *testing.T) {
	const input = `---
unknown: field
...

# Recipe`
	p := NewParser()

	meta, err := p.ParseRecipe([]byte(input))
	assert.Error(t, err)
	assert.Nil(t, meta)
}
