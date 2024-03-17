import type Fraction from 'fraction.js';

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
		text: string;
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

export interface Metadata {
	readonly servings: number;
	readonly tags: string[];
	readonly source?: string;
}

export interface Recipe {
	readonly metadata: Metadata;
	readonly title: string;

	readonly intro: Section;
	readonly steps: Section[];
}
