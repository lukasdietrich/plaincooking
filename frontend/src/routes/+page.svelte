<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<ul class="grid grid-cols-3 gap-5">
	{#each data.recipes as recipe}
		<li class="rounded-sm shadow">
			<a class="block p-3" href="/recipes/{recipe.id}">
				{recipe.title}
			</a>
		</li>
	{/each}
</ul>
