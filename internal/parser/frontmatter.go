package parser

import (
	"bytes"
	"slices"

	"gopkg.in/yaml.v3"
)

var (
	fenceStart = "---"
	fenceEnd   = "..."
)

type frontmatter[M any] struct {
	Matter  *M
	Content []byte
}

func newFrontmatter[M any](matterBytes []byte, contentBytes []byte) (*frontmatter[M], error) {
	if matterBytes == nil {
		return &frontmatter[M]{Content: contentBytes}, nil
	}

	decoder := yaml.NewDecoder(bytes.NewReader(matterBytes))
	decoder.KnownFields(true)

	var matter M
	if err := decoder.Decode(&matter); err != nil {
		return nil, err
	}

	return &frontmatter[M]{Matter: &matter, Content: contentBytes}, nil
}

func parseFrontmatter[M any](text []byte) (*frontmatter[M], error) {
	lines := splitLines(text)

	if len(lines) > 0 && lines[0] == fenceStart {
		if end := slices.Index(lines, fenceEnd); end > -1 {
			matterBytes := joinLines(lines[1:end])
			contentBytes := joinLines(lines[end+1:])

			return newFrontmatter[M](matterBytes, contentBytes)
		}
	}

	return newFrontmatter[M](nil, text)
}
