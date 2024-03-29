<script lang="ts">
	import type { components } from '$lib/api';
	import { createEventDispatcher } from 'svelte';
	import { scale } from 'svelte/transition';

	export let image: components['schemas']['AssetMetadata'];

	const dispatch = createEventDispatcher();

	function handleClose() {
		dispatch('close');
	}
</script>

<button
	class="fixed inset-0 z-750 bg-black/90 flex items-center justify-center"
	on:click={handleClose}
	transition:scale={{ start: 1.05, duration: 200, delay: 150 }}
>
	<div class="w-full h-full p-15 rounded-10">
		<img src={image.href} alt={image.id} class="mx-auto h-full object-scale-down" />
	</div>
</button>
