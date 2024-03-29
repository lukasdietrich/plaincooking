import type { components } from '$lib/api';
import { registerMiddleware } from '$lib/api';
import { notify } from './store';

export { notify };
export { default as Notifications } from './Notifications.svelte';

registerMiddleware({
	async onResponse(res: Response) {
		if (res.status >= 300) {
			const error = <components['schemas']['PlaincookingApiError']>await res.clone().json();

			notify({
				type: 'error',
				text: `api.error.${error.code}`
			});
		}

		return undefined;
	}
});
