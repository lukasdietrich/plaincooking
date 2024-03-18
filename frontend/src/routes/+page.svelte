<script lang="ts">
	import type { PageData } from './$types';
	import placeholder from '$lib/images/placeholder.jpeg';
	import { t } from '$lib/i18n';
	import { ActionPortal, Action } from '$lib/components/actions';

	export let data: PageData;
</script>

<ActionPortal>
	<Action href={'/recipes/new'} title={$t('actions.recipe.new')}>
		<i class="icon-plus"></i>
	</Action>
</ActionPortal>

<ul class="grid md:grid-cols-2 lg:grid-cols-3 gap-5 group">
	{#each data.recipes as recipe}
		<li
			class="rounded shadow overflow-clip bg-cover transition group-hover:opacity-80 hover:(!opacity-100 ring-2 ring-emerald-800)"
			style:background-image="url({placeholder})"
		>
			<a class="block" href="/recipes/{recipe.id}">
				<div class="flex flex-col justify-end h-64">
					<span class="rounded bg-black/80 text-white text-center font-semibold p-3 m-1">
						{recipe.title}
					</span>
				</div>
			</a>
		</li>
	{/each}
</ul>
