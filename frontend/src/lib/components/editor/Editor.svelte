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

<div class="editor bg-gray-50 font-mono text-lg" bind:this={element}></div>

<style class="css">
	.editor :global(.cm-heading-1 *) {
		--at-apply: text-2xl thickness-1 underline-offset-4; /**/
	}

	.editor :global(.cm-frontmatter-fence *) {
		--at-apply: text-gray-900/40 font-semibold; /**/
	}

	.editor :global(.cm-hr) {
		--at-apply: bg-blue-100; /**/
	}

	.editor :global(.cm-hr *) {
		--at-apply: text-blue-900/40 font-semibold; /**/
	}

	.editor :global(.cm-ingredient) {
		--at-apply: bg-emerald-100 font-semibold not-italic; /**/
	}

	.editor :global(.cm-codemark *) {
		--at-apply: text-emerald-900/40; /**/
	}
</style>
