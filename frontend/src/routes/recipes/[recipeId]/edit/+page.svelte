<script lang="ts">
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import { resolveRoute } from '$app/paths';
	import { page } from '$app/stores';
	import { client } from '$lib';
	import { onMount } from 'svelte';
	import { basicSetup, EditorView } from 'codemirror';
	import { markdown } from '@codemirror/lang-markdown';
	import { ActionPortal, Action } from '$lib/components/actions';

	export let data: PageData;

	let editorHost!: HTMLElement;
	let editor: EditorView;

	onMount(() => {
		const theme = EditorView.theme({
			'.cm-content': {
				fontFamily: 'inherit',
				fontSize: 'inherit'
			},

			'&.cm-focused': {
				outline: 'none'
			}
		});

		editor = new EditorView({
			doc: data.content,
			extensions: [basicSetup, theme, EditorView.lineWrapping, markdown()],
			parent: editorHost
		});
	});

	async function save() {
		const content = editor.state.doc.toString();
		const recipeId = $page.params.recipeId;

		await client.recipes.updateRecipe(recipeId, content);
		await goto(resolveRoute('/recipes/[recipeId]', $page.params), { invalidateAll: true });
	}
</script>

<ActionPortal>
	<Action href={resolveRoute('/recipes/[recipeId]', $page.params)}>
		<i class="icon-undo-2"></i>
	</Action>

	<Action on:click={save}>
		<i class="icon-save"></i>
	</Action>
</ActionPortal>

<div class="bg-gray-50 font-mono text-lg" bind:this={editorHost}></div>
