import type { paths } from './types.gen';
import type { MediaType } from 'openapi-typescript-helpers';
import type { ClientOptions, FetchResponse, BodySerializer } from 'openapi-fetch';
import createClient from 'openapi-fetch';
import { middleware } from './middleware';

const jsonSerializer: BodySerializer<never> = (body) => JSON.stringify(body);
const textSerializer: BodySerializer<never> = (body) => body;

const options: ClientOptions = {
	baseUrl: '/api',
	bodySerializer: jsonSerializer
};

export function createApi(fetch?: typeof globalThis.fetch, skipMiddleware?: boolean) {
	const client = createClient<paths>({ ...options, fetch });
	if (skipMiddleware !== true) {
		client.use(...middleware());
	}

	const { GET, POST, PUT, DELETE } = client;

	return {
		readUserInfo() {
			return handleResponse(GET('/session/info'));
		},

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
		},

		uploadRecipeImage(recipeId: string, image: File) {
			return handleResponse(
				POST('/recipes/{recipeId}/images', {
					params: {
						path: {
							recipeId
						}
					},
					body: {
						image
					},
					bodySerializer({ image }) {
						const formData = new FormData();
						formData.append('image', image);
						return formData;
					}
				})
			);
		}
	};
}

async function handleResponse<
	T extends Record<string | number, any>,
	Options,
	Media extends MediaType
>(responsePromise: Promise<FetchResponse<T, Options, Media>>) {
	const { error, data } = await responsePromise;

	if (error) {
		throw error;
	}

	return data!;
}
