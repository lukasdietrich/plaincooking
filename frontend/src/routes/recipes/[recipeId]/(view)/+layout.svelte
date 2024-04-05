<script lang="ts">
	import type { LayoutData } from './$types';
	import { page } from '$app/stores';
	import { resolveRoute } from '$lib/routing';
	import { t } from '$lib/i18n';
	import { ActionPortal, Action } from '$lib/components/actions';
	import { Header } from '$lib/components/recipe';
	import { Gallery } from '$lib/components/gallery';

	export let data: LayoutData;

	$: recipeId = $page.params.recipeId;
	$: recipe = data.recipe;
	$: images = data.images;
</script>

<ActionPortal>
	<Action
		href={resolveRoute('/recipes/[recipeId]/edit', $page)}
		title={$t('actions.recipe.edit')}
		requireAuth={true}
	>
		<i class="icon-file-pen-line"></i>
	</Action>
</ActionPortal>

<article class="flex flex-col space-y-5">
	<Gallery {recipeId} {images} />
	<Header {recipe} />
	<slot />
</article>
