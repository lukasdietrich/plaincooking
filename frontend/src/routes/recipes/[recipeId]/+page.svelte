<script lang="ts">
	import type { PageData } from './$types';
	import { page } from '$app/stores';
	import { resolveRoute } from '$app/paths';
	import { Renderer } from '$lib/components/recipe/renderer';
	import ActionPortal from '$lib/components/actions/ActionPortal.svelte';
	import Action from '$lib/components/actions/Action.svelte';

	export let data: PageData;
</script>

<ActionPortal>
	<Action href={resolveRoute('/', $page.params)}>
		<i class="icon-undo-2"></i>
	</Action>

	<Action href={resolveRoute('/recipes/[recipeId]/edit', $page.params)}>
		<i class="icon-file-pen-line"></i>
	</Action>
</ActionPortal>

<Renderer content={data.content}></Renderer>
