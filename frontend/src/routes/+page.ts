import type { PageLoad } from './$types';
import { client } from '$lib';

export const load: PageLoad = async () => {
	return {
		recipes: await client.recipes.listRecipes()
	};
};
