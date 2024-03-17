package parser

import (
	"errors"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var (
	ErrInvalidRecipe = errors.New("invalid recipe")
	ErrMissingTitle  = fmt.Errorf("%w: missing title", ErrInvalidRecipe)
)

type RecipeFrontmatter struct {
	Servings uint
	Tags     []string
	Source   string
}

type RecipeMetadata struct {
	*RecipeFrontmatter
	Title string
}

type RecipeParser struct {
	m goldmark.Markdown
}

func NewParser() *RecipeParser {
	m := goldmark.New()

	return &RecipeParser{m: m}
}

func (p *RecipeParser) ParseRecipe(content []byte) (*RecipeMetadata, error) {
	frontmatter, err := parseFrontmatter[RecipeFrontmatter](content)
	if err != nil {
		return nil, err
	}

	r := text.NewReader(frontmatter.Content)
	ctx := parser.NewContext()
	root := p.m.Parser().Parse(r, parser.WithContext(ctx))

	recipe := RecipeMetadata{RecipeFrontmatter: frontmatter.Matter}

	if err := p.findTitle(&recipe, frontmatter.Content, root); err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (p *RecipeParser) findTitle(recipe *RecipeMetadata, content []byte, node ast.Node) error {
	err := ast.Walk(node, func(n ast.Node, _ bool) (ast.WalkStatus, error) {
		if heading, ok := n.(*ast.Heading); ok && heading.Level == 1 {
			recipe.Title = string(n.Text(content))
			return ast.WalkStop, nil
		}

		if _, ok := n.(*ast.ThematicBreak); ok {
			return ast.WalkStop, nil
		}

		return ast.WalkContinue, nil
	})

	if err != nil {
		return err
	}

	if recipe.Title == "" {
		return ErrMissingTitle
	}

	return nil
}
