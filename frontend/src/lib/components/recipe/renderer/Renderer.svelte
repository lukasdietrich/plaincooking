<script lang="ts">
	import { parse } from './parser';
	import MarkdownToken from './MarkdownToken.svelte';
	import Step from './Step.svelte';

	export let content: string;

	$: recipe = parse(content);
</script>

<article class="flex flex-col space-y-5">
	<header>
		<MarkdownToken token={recipe.intro} />
	</header>

	{#each recipe.steps as tokens, index}
		<Step {index} {tokens} />
	{/each}
</article>
