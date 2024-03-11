<script lang="ts">
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { basicSetup, EditorView } from 'codemirror';
	import { markdown } from '@codemirror/lang-markdown';

	export let data: PageData;

	let editorHost!: HTMLElement;

	onMount(() => {
		console.log(editorHost);

		new EditorView({
			doc: data.content,
			extensions: [basicSetup, markdown()],
			parent: editorHost
		});
	});
</script>

<div bind:this={editorHost}></div>
