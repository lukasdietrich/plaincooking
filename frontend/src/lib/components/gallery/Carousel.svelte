<script lang="ts">
	import type { components } from '$lib/api';
	import { createEventDispatcher } from 'svelte';
	import { user } from '$lib/auth';
	import { dragOnBody, dropOnBody } from '$lib/actions/drag';

	export let images: components['schemas']['AssetMetadata'][] = [];

	const dispatch = createEventDispatcher();

	let current = -1;
	$: current = getInitialCurrent(images);

	function getInitialCurrent<T>(images: T[]) {
		if (images.length > 0) {
			return 0;
		}

		return -1;
	}

	function handleSpotlight(image: components['schemas']['AssetMetadata']) {
		dispatch('spotlight', image);
	}

	function handleFileChange({ target }: Event) {
		const input = <HTMLInputElement>target;
		const file = input.files?.item(0);

		if (file) {
			dispatch('upload', file);
		}
	}

	function handleFileDrop({ detail }: CustomEvent<File>) {
		dispatch('upload', detail);
	}
</script>

<div class="h-full text-white text-lg overflow-clip">
	<div class="relative h-full transition-750" style:transform="translateX({current * -101}%)">
		<div class="absolute w-full h-full -left-101% bg-slate-50 border-2 border-slate-200 rounded">
			<label
				for="image-upload"
				class="flex items-center justify-center w-full h-full cursor-pointer"
			>
				<i class="icon-image-up text-3xl text-slate-700/70"></i>
			</label>

			{#if $user}
				<input
					id="image-upload"
					type="file"
					class="hidden"
					accept="image/png, image/jpeg"
					on:change={handleFileChange}
					on:dragEnterBody={() => (current = -1)}
					on:dragLeaveBody={() => (current = getInitialCurrent(images))}
					on:dropBody={handleFileDrop}
					use:dragOnBody
					use:dropOnBody
				/>
			{/if}

			{#if images.length > 0}
				<button class="right-0 rounded-r arrow" on:click={() => (current = 0)}>
					<i class="icon-arrow-right"></i>
				</button>
			{/if}
		</div>

		{#each images as image, index}
			<div class="absolute w-full h-full rounded overflow-clip" style:left="{index * 101}%">
				<button
					class="block w-full h-full bg-cover bg-center transition-2000"
					class:hover:scale-105={index === current}
					class:filter-blur-2={index !== current}
					style:background-image="url({image.href}?thumbnail=banner)"
					on:click={() => handleSpotlight(image)}
				/>

				{#if index > 0 || $user}
					<button class="left-0 rounded-l arrow" on:click={() => (current = index - 1)}>
						<i class:icon-arrow-left={index > 0} class:icon-image-up={index === 0}></i>
					</button>
				{/if}

				{#if index < images.length - 1}
					<button class="right-0 rounded-r arrow" on:click={() => (current = index + 1)}>
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
