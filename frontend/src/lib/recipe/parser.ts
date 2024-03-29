import type { Token as MarkedToken } from 'marked';
import type { Recipe, Metadata, Token, Tokens } from './model';
import Fraction from 'fraction.js';
import { Marked } from 'marked';
import { parseFrontmatter } from './frontmatter';

const marked = new Marked();

export function parseRecipe(src: string): Recipe {
	const { matter, content } = parseFrontmatter<Metadata>(src);
	const tokens = processTokenArray(marked.lexer(content));
	const [intro, ...steps] = splitChunks(tokens, isHorizontalRule);
	const title = findTitle(intro);

	if (!title) {
		throw new Error('missing title');
	}

	const metadata = {
		servings: 1,
		tags: [],

		...matter
	};

	return { metadata, title, intro, steps };
}

function isHorizontalRule({ type }: Token): boolean {
	return type === 'hr';
}

function splitChunks<E>(arr: E[], separator: (element: E) => boolean): E[][] {
	const chunks: E[][] = [];
	let chunk: E[] = [];

	for (const element of arr) {
		if (separator(element)) {
			chunks.push(chunk);
			chunk = [];
		} else {
			chunk.push(element);
		}
	}

	if (chunk.length > 0) {
		chunks.push(chunk);
	}

	return chunks;
}

function processTokenArray(tokens?: MarkedToken[]): Token[] {
	if (!tokens) {
		return [];
	}

	return <Token[]>tokens.map(processToken).filter((t) => t !== undefined);
}

function processToken(token: MarkedToken): Token | undefined {
	switch (token.type) {
		case 'codespan':
			return parseIngredient(token.text);

		case 'em':
			return {
				type: 'em',
				children: processTokenArray(token.tokens)
			};

		case 'heading':
			return {
				type: 'heading',
				level: token.depth,
				text: token.text,
				children: processTokenArray(token.tokens)
			};

		case 'hr':
			return {
				type: 'hr'
			};

		case 'paragraph':
			return {
				type: 'paragraph',
				children: processTokenArray(token.tokens)
			};

		case 'strong':
			return {
				type: 'strong',
				children: processTokenArray(token.tokens)
			};

		case 'text':
			return {
				type: 'text',
				text: token.text
			};

		default:
			return undefined;
	}
}

function parseIngredient(text: string): Tokens.Ingredient {
	const pattern = /((?<quantity>[0-9/.]+)(?<unit>[\w]+)? )?(?<ingredient>.*)/;
	const match = text.match(pattern);

	if (!match) {
		return {
			type: 'ingredient',
			ingredient: text
		};
	}

	const { quantity, unit, ingredient } = <{ quantity?: string; unit?: string; ingredient: string }>(
		match.groups
	);

	return {
		type: 'ingredient',
		quantity: quantity?.length ? new Fraction(quantity) : undefined,
		unit,
		ingredient
	};
}

function findTitle(tokens: Token[]): string | undefined {
	for (const token of tokens) {
		if (token.type === 'heading' && token.level === 1) {
			return token.text;
		}
	}

	return undefined;
}
