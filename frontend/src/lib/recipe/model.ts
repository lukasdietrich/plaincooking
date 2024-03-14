import type { Fraction } from 'mathjs';

export namespace Tokens {
	interface WithChildren {
		children: Token[];
	}

	export interface Em extends WithChildren {
		type: 'em';
	}

	export interface Heading extends WithChildren {
		type: 'heading';
		level: number;
	}

	export interface Hr {
		type: 'hr';
	}

	export interface Ingredient {
		type: 'ingredient';
		quantity?: Fraction;
		unit?: string;
		ingredient: string;
	}

	export interface Paragraph extends WithChildren {
		type: 'paragraph';
	}

	export interface Strong extends WithChildren {
		type: 'strong';
	}

	export interface Text {
		type: 'text';
		text: string;
	}
}

export type Token =
	| Tokens.Em
	| Tokens.Heading
	| Tokens.Hr
	| Tokens.Ingredient
	| Tokens.Paragraph
	| Tokens.Strong
	| Tokens.Text;
export type Section = Token[];

interface Metadata {
	servings: number;
	tags: string[];
	source?: string;
}

export interface Recipe {
	metadata: Metadata;
	intro: Section;
	steps: Section[];
}
