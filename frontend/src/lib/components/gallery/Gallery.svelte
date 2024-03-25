<script lang="ts">
	import type { components } from '$lib/api';
	import Carousel from './Carousel.svelte';
	import Spotlight from './Spotlight.svelte';

	export let images: components['schemas']['AssetMetadata'][] = [];

	let spotlight: components['schemas']['AssetMetadata'] | null = null;
</script>

<div class="relative h-36 sm:h-48 md:h-72">
	<Carousel {images} on:spotlight={(event) => (spotlight = event.detail)} />

	{#if spotlight}
		<Spotlight image={spotlight} on:close={() => (spotlight = null)} />
	{/if}
</div>
