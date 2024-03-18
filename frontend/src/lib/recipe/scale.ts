import type { Recipe, Token } from './model';
import Fraction from 'fraction.js';

export function scaleRecipe({ metadata, title, intro, steps }: Recipe, servings: number): Recipe {
	const multiplier = new Fraction(servings, metadata.servings);

	return {
		metadata: {
			...metadata,
			servings
		},
		title,
		intro: scaleTokens(intro, multiplier),
		steps: steps.map((step) => scaleTokens(step, multiplier))
	};
}

function scaleTokens(tokens: Token[], multiplier: Fraction): Token[] {
	return tokens.map((token) => scaleToken(token, multiplier));
}

function scaleToken(token: Token, multiplier: Fraction): Token {
	if (token.type === 'ingredient' && token.quantity) {
		return {
			...token,
			quantity: token.quantity.mul(multiplier)
		};
	}

	if ('children' in token) {
		return {
			...token,
			children: scaleTokens(token.children, multiplier)
		};
	}

	return token;
}
