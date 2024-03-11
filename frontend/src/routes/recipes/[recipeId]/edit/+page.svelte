<script lang="ts">
	import type { PageData } from './$types';
	import { page } from '$app/stores';
	import { client } from '$lib';
	import { onMount } from 'svelte';
	import { basicSetup, EditorView } from 'codemirror';
	import { markdown } from '@codemirror/lang-markdown';

	export let data: PageData;

	let editorHost!: HTMLElement;
	let editor: EditorView;

	onMount(() => {
		editor = new EditorView({
			doc: data.content,
			extensions: [basicSetup, markdown()],
			parent: editorHost
		});
	});

	async function save() {
		const content = editor.state.doc.toString();
		const recipeId = $page.params.recipeId;

		await client.recipes.updateRecipe(recipeId, content);
	}
</script>

<div bind:this={editorHost}></div>
<button on:click={save}>Save</button>
