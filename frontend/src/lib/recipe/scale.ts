import type { Recipe, Token } from './model';
import Fraction from 'fraction.js';

export function scaleRecipe({ metadata, title, intro, steps }: Recipe, servings: number): Recipe {
	const multiplier = new Fraction(servings, metadata.servings);
	const scale = scaleTokens(multiplier);

	return {
		metadata: {
			...metadata,
			servings
		},
		title,
		intro: scale(intro),
		steps: steps.map(scale)
	};
}

const scaleTokens =
	(multiplier: Fraction) =>
	(tokens: Token[]): Token[] =>
		tokens.map(scaleToken(multiplier));

const scaleToken =
	(multiplier: Fraction) =>
	(token: Token): Token => {
		if (token.type === 'ingredient' && token.quantity) {
			return {
				...token,
				quantity: token.quantity.mul(multiplier)
			};
		}

		if ('children' in token) {
			return {
				...token,
				children: scaleTokens(multiplier)(token.children)
			};
		}

		return token;
	};
