import type { paths } from './types.gen';
import type { ClientOptions, FetchResponse, BodySerializer } from 'openapi-fetch';
import createClient from 'openapi-fetch';

export * from './types.gen';

const jsonSerializer: BodySerializer<never> = (body) => JSON.stringify(body);
const textSerializer: BodySerializer<never> = (body) => body;

const options: ClientOptions = {
	baseUrl: '/api',
	bodySerializer: jsonSerializer
};

export function createApi(fetch?: typeof globalThis.fetch) {
	const { GET, POST, PUT, DELETE } = createClient<paths>({ ...options, fetch });

	return {
		listRecipes() {
			return handleResponse(GET('/recipes'));
		},

		readRecipe(recipeId: string) {
			return handleResponse(
				GET('/recipes/{recipeId}', {
					params: {
						path: {
							recipeId
						}
					},
					parseAs: 'text'
				})
			);
		},

		createRecipe(content: string) {
			return handleResponse(
				POST('/recipes', {
					headers: {
						'Content-Type': 'text/markdown'
					},
					body: content,
					bodySerializer: textSerializer
				})
			);
		},

		updateRecipe(recipeId: string, content: string) {
			return handleResponse(
				PUT('/recipes/{recipeId}', {
					headers: {
						'Content-Type': 'text/markdown'
					},
					params: {
						path: {
							recipeId
						}
					},
					body: content,
					bodySerializer: textSerializer
				})
			);
		},

		deleteRecipe(recipeId: string) {
			return handleResponse(
				DELETE('/recipes/{recipeId}', {
					params: {
						path: {
							recipeId
						}
					}
				})
			);
		},

		listRecipeImages(recipeId: string) {
			return handleResponse(
				GET('/recipes/{recipeId}/images', {
					params: {
						path: {
							recipeId
						}
					}
				})
			);
		}
	};
}

async function handleResponse<T, O>(responsePromise: Promise<FetchResponse<T, O>>) {
	const { error, data } = await responsePromise;

	if (error) {
		throw error;
	}

	return data!;
}
