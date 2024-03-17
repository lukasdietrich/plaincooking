import type { LayoutLoad } from './$types';
import { client } from '$lib';
import { parse } from '$lib/recipe';

export const load: LayoutLoad = async ({ params }) => {
	const [metadata, content] = await Promise.all([
		client.recipes.readRecipeMetadata(params.recipeId),
		client.recipes.readRecipe(params.recipeId)
	]);

	return {
		title: metadata.title,

		metadata,
		content,
		recipe: parse(content)
	};
};
