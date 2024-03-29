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
		class="fixed inset-0 z-750 bg-black/80 backdrop-blur-30 backdrop-grayscale-70"
		on:click={handleClose}
		transition:scale|global={{ start: 1.05, duration: 300 }}
	>
		<div class="flex w-full h-full p-3 md:p-5 lg:p-15">
			<img
				src={image.src}
				alt={image.alt}
				class="m-auto h-auto max-h-full object-scale-down rounded-lg"
			/>
		</div>
	</button>
{/await}
