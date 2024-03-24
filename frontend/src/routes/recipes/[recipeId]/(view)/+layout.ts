import type { LayoutLoad } from './$types';
import { createApi } from '$lib/api';
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

export const load: LayoutLoad = async ({ fetch, parent, params, url }) => {
	const { listRecipeImages } = createApi(fetch);

	const { recipeId } = params;
	const { recipe } = await parent();
	const images = await listRecipeImages(recipeId);
	const servings = parseServings(url) ?? recipe.metadata.servings;

	return {
		recipe: scale(recipe, servings),
		images
	};
};
