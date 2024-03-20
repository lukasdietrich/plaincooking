import type { LayoutLoad } from './$types';
import { createApi } from '$lib/api';
import { parse } from '$lib/recipe';

export const load: LayoutLoad = async ({ fetch, params }) => {
	const { readRecipe } = createApi(fetch);
	const content = await readRecipe(params.recipeId);
	const recipe = parse(content);

	return {
		title: recipe.title,

		content,
		recipe
	};
};
