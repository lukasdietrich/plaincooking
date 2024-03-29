<script lang="ts">
	import type { components } from '$lib/api';
	import { createEventDispatcher } from 'svelte';
	import { scale } from 'svelte/transition';

	export let image: components['schemas']['AssetMetadata'];

	const dispatch = createEventDispatcher();

	$: preload = new Promise<HTMLImageElement>((resolve) => {
		const preload = new Image();

		preload.src = image.href;
		preload.addEventListener('load', () => resolve(preload));
	});

	function handleClose() {
		dispatch('close');
	}
</script>

{#await preload then image}
	<button
		class="fixed inset-0 z-750 bg-black/90 flex items-center justify-center"
		on:click={handleClose}
		transition:scale={{ start: 1.05, duration: 200, delay: 150 }}
	>
		<div class="w-full h-full p-15 rounded-10">
			<img src={image.src} alt={image.alt} class="mx-auto h-full object-scale-down" />
		</div>
	</button>
{/await}
