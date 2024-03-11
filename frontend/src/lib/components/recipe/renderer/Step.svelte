<script lang="ts">
	import type { Token } from 'marked';
	import MarkdownToken from './MarkdownToken.svelte';

	export let index: number = 0;
	export let tokens: Token[] = [];

	let done = false;

	$: id = `step-${index}`;
</script>

<section class="ring ring-emerald-600 rounded">
	<div class="flex select-none bg-emerald-100 rounded" class:line-through={done}>
		<label class="w-full px-3 py-2 font-semibold text-emerald-700" for={id}>
			Step {index + 1}
		</label>
		<input class="mx-3" type="checkbox" {id} bind:checked={done} />
	</div>

	<div class="px-3 py-2" class:hidden={done}>
		<MarkdownToken token={tokens} />
	</div>
</section>
