import type { LayoutLoad } from './$types';
import { client } from '$lib';
import { parse } from '$lib/recipe';

export const load: LayoutLoad = async ({ params }) => {
	const content = await client.recipes.readRecipe(params.recipeId);
	const recipe = parse(content);

	return {
		title: recipe.title,

		content,
		recipe
	};
};
