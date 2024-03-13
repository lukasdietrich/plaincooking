<script lang="ts">
	import type { PageData } from './$types';
	import { page } from '$app/stores';
	import { resolveRoute } from '$app/paths';
	import { parse } from '$lib/recipe';
	import { Recipe } from '$lib/components/recipe';
	import { ActionPortal, Action } from '$lib/components/actions';

	export let data: PageData;

	$: recipe = parse(data.content);
</script>

<ActionPortal>
	<Action href={resolveRoute('/', $page.params)}>
		<i class="icon-undo-2"></i>
	</Action>

	<Action href={resolveRoute('/recipes/[recipeId]/edit', $page.params)}>
		<i class="icon-file-pen-line"></i>
	</Action>
</ActionPortal>

<Recipe {recipe} />
