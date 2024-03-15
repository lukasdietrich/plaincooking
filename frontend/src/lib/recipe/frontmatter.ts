import { parse as parseYaml } from 'yaml';

export const FENCE_START = '---';
export const FENCE_END = '...';

export interface Frontmatter<M> {
	matter?: M;
	content: string;
}

export function parseFrontmatter<M>(text: string): Frontmatter<M> {
	const lines = splitLines(text);
	const [first, ...rest] = lines;

	if (first === FENCE_START) {
		const end = rest.findIndex((line) => line === FENCE_END);
		if (end > -1) {
			return {
				matter: parseYaml(joinLines(rest.slice(0, end))),
				content: joinLines(rest.slice(end + 1))
			};
		}
	}

	return { content: text };
}

function splitLines(text: string): string[] {
	return text.split('\n');
}

function joinLines(lines: string[]): string {
	return lines.join('\n');
}
