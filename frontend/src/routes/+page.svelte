<script lang="ts">
	import type { PageData } from './$types';
	import placeholder from '$lib/images/placeholder.jpeg';

	export let data: PageData;
</script>

<ul class="grid grid-cols-3 gap-5 group">
	{#each data.recipes as recipe}
		<li
			class="rounded shadow overflow-clip bg-cover transition group-hover:opacity-80 hover:(!opacity-100 ring-2 ring-emerald-800)"
			style:background-image="url({placeholder})"
		>
			<a class="block" href="/recipes/{recipe.id}">
				<div class="flex flex-col justify-end h-64">
					<span class="bg-black/70 text-white font-semibold p-3 m-2">
						{recipe.title}
					</span>
				</div>
			</a>
		</li>
	{/each}
</ul>
