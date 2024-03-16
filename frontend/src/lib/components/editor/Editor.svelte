<script lang="ts">
	import type { ViewUpdate } from '@codemirror/view';
	import { EditorView } from '@codemirror/view';
	import { onMount, onDestroy } from 'svelte';
	import { create } from './codemirror';

	export let value: string = '';

	let element: HTMLElement;
	let editor: EditorView;

	function handleUpdate(update: ViewUpdate) {
		if (update.docChanged) {
			const doc = update.state.doc;
			value = doc.toString();
		}
	}

	onMount(() => {
		editor = create(element, value, [EditorView.updateListener.of(handleUpdate)]);
	});

	onDestroy(() => {
		editor?.destroy();
	});
</script>

<div bind:this={element} />
