<script lang="ts">
	import type { PageData } from './$types';
	import { Renderer } from '$lib/components/recipe/renderer';

	export let data: PageData;
</script>

<Renderer content={data.content}></Renderer>
