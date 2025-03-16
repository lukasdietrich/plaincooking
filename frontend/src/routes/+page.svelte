<script lang="ts">
	import type { PageData } from './$types';
	import type { components } from '$lib/api';
	import placeholder from '$lib/images/placeholder.jpeg';
	import { t } from '$lib/i18n';
	import { ActionPortal, Action } from '$lib/components/actions';

	export let data: PageData;

	function getImageUrl(recipe: components['schemas']['RecipeMetadata']) {
		if (recipe.imageHref) {
			return `${recipe.imageHref}?thumbnail=tile`;
		}

		return placeholder;
	}
</script>

<ActionPortal>
	<Action href="/recipes/new" title={$t('actions.recipe.new')} requireAuth={true}>
		<i class="icon-plus"></i>
	</Action>
</ActionPortal>

<ul class="grid md:grid-cols-2 lg:grid-cols-3 gap-5 group">
	{#each data.recipes as recipe (recipe.id)}
		<li
			class="rounded shadow overflow-clip bg-center bg-cover transition group-hover:opacity-80 hover:(!opacity-100 ring-2 ring-emerald-800)"
			style:background-image="url({getImageUrl(recipe)})"
		>
			<a class="block" href="/recipes/{recipe.id}">
				<div class="flex flex-col justify-end h-64">
					<span
						class="bg-black/30 backdrop-sepia-30 backdrop-blur-5 text-white text-shadow-sm text-center font-semibold p-3"
					>
						{recipe.title}
					</span>
				</div>
			</a>
		</li>
	{/each}
</ul>
