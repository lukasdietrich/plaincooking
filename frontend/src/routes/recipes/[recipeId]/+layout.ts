import type { PageLoad } from './$types';
import { client } from '$lib';

export const load: PageLoad = async ({ params }) => {
	return {
		content: await client.recipes.readRecipe(params.recipeId)
	};
};
