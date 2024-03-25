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
	class="absolute -top-3 -left-3 -right-3 shadow-xl shadow-black/15 rounded overflow-clip transition-300"
	class:opacity-0={hidden}
>
	<img src={image.href} alt={image.id} on:load={() => (loading = false)} />

	<button class="absolute right-3 top-3 rounded text-white text-lg" on:click={handleClose}>
		<i class="icon-x bg-black/30 backdrop-blur-10 p-2 rounded-full transition hover:bg-red-600/50"
		></i>
	</button>
</div>
