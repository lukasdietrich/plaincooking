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
