<script lang="ts">
	import type { components } from '$lib/api';
	import { createApi } from '$lib/api';
	import Carousel from './Carousel.svelte';
	import Spotlight from './Spotlight.svelte';

	export let recipeId: string;
	export let images: components['schemas']['AssetMetadata'][] = [];

	let imagesWithNew = images;
	let spotlight: components['schemas']['AssetMetadata'] | null = null;

	$: imagesWithNew = images;

	async function handleUpload(file: File) {
		const { uploadRecipeImage } = createApi();
		const image = await uploadRecipeImage(recipeId, file);

		imagesWithNew = [image, ...imagesWithNew];
	}
</script>

<div class="relative h-36 sm:h-48 md:h-72">
	<Carousel
		images={imagesWithNew}
		on:spotlight={(event) => (spotlight = event.detail)}
		on:upload={(event) => handleUpload(event.detail)}
	/>

	{#if spotlight}
		<Spotlight image={spotlight} on:close={() => (spotlight = null)} />
	{/if}
</div>
