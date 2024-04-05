<script lang="ts">
	import type { components } from '$lib/api';

	export let user: components['schemas']['UserInfo'];
</script>

<div class="h-full px-5 py-2 bg-emerald-700">
	<img
		class="h-full object-fit rounded-full ring-2 ring-emerald-900"
		src={user.pictureUrl}
		alt={user.name}
	/>
</div>
