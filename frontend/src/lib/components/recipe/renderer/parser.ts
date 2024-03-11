import type { Token } from 'marked';
import { Marked } from 'marked';

const marked = new Marked();

export interface Recipe {
	intro: Token[];
	steps: Token[][];
}

export function parse(src: string): Recipe {
	const tokens = marked.lexer(src);
	const [intro, ...steps] = splitChunks(tokens, isHorizontalRule);

	return { intro, steps };
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
