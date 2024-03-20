import type { PageLoad } from './$types';
import { createApi } from '$lib/api';

export const load: PageLoad = async ({ fetch }) => {
	const { listRecipes } = createApi(fetch);
	const recipes = await listRecipes();

	return {
		recipes
	};
};
