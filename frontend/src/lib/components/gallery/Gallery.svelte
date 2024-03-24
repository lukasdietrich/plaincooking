<script lang="ts">
	import type { components } from '$lib/api';
	import Carousel from './Carousel.svelte';

	export let images: components['schemas']['AssetMetadataDto'][] = [];

	$: hrefs = images.map((image) => image.href);
</script>

<div class="relative h-36 sm:h-48 md:h-72">
	<Carousel images={hrefs} />
</div>
