import type { LayoutLoad } from './$types';
import { client } from '$lib';

export const load: LayoutLoad = async ({ params }) => {
	const [metadata, content] = await Promise.all([
		client.recipes.readRecipeMetadata(params.recipeId),
		client.recipes.readRecipe(params.recipeId)
	]);

	return {
		title: metadata.title,

		metadata,
		content
	};
};
