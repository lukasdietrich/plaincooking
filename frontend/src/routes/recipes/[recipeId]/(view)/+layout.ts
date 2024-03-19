import type { LayoutLoad } from './$types';
import { scale } from '$lib/recipe';

function parseServings(url: URL): number | null {
	const query = url.searchParams.get('servings');
	if (query) {
		const servings = parseInt(query);
		if (!isNaN(servings)) {
			return Math.max(1, servings);
		}
	}

	return null;
}

export const load: LayoutLoad = async ({ parent, url }) => {
	const { recipe } = await parent();
	const servings = parseServings(url) ?? recipe.metadata.servings;

	return {
		recipe: scale(recipe, servings)
	};
};
