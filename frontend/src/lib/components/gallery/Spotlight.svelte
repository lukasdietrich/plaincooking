<script lang="ts">
	import type { components } from '$lib/api';
	import { onMount, createEventDispatcher } from 'svelte';

	export let image: components['schemas']['AssetMetadata'];

	const dispatch = createEventDispatcher();

	let loading = true;
	let mounted = false;

	onMount(() => setTimeout(() => (mounted = true), 10));

	$: hidden = loading || !mounted;

	function handleClose() {
		dispatch('close');
	}
</script>

<div
	class="fixed inset-0 z-750 transition-300 bg-black/90 flex items-center justify-center"
	class:opacity-0={hidden}
>
	<div class="w-full h-full p-15 rounded-10">
		<img
			src={image.href}
			alt={image.id}
			on:load={() => (loading = false)}
			class="mx-auto h-full object-scale-down"
		/>
	</div>

	<button class="absolute right-5 top-5 rounded text-white text-lg" on:click={handleClose}>
		<i class="icon-x bg-white/30 backdrop-blur-10 p-2 rounded-full transition hover:bg-red-700/80"
		></i>
	</button>
</div>
