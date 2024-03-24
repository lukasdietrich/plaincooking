<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let images: string[] = [];

	const dispatch = createEventDispatcher();

	let current = 0;

	function handleExpand() {
		dispatch('expand');
	}
</script>

<div class="h-full text-white text-lg overflow-clip">
	<div class="relative h-full transition-750" style:transform="translateX(-{current * 101}%)">
		{#each images as image, index}
			<div class="absolute w-full h-full rounded overflow-clip" style:left="{index * 101}%">
				<button
					class="block w-full h-full bg-cover bg-center transition-2000"
					class:hover:scale-105={index === current}
					class:filter-blur-2={index !== current}
					style:background-image="url({image}?thumbnail=banner)"
					on:click={handleExpand}
				/>

				{#if index > 0}
					<button class="left-0 rounded-l arrow" on:click={() => current--}>
						<i class="icon-arrow-left"></i>
					</button>
				{/if}

				{#if index < images.length - 1}
					<button class="right-0 rounded-r arrow" on:click={() => current++}>
						<i class="icon-arrow-right"></i>
					</button>
				{/if}
			</div>
		{/each}
	</div>
</div>

<style>
	.arrow {
		--at-apply: 'absolute top-0 bottom-0 p-3 transition hover:(backdrop-blur-5 bg-black/10)';
	}

	.arrow i {
		--at-apply: 'bg-black/30 backdrop-blur-10 p-2 rounded-full';
	}
</style>
