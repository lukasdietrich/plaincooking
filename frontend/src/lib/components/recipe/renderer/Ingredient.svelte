<script lang="ts">
	export let text: string;

	const id = Math.random().toString(16);

	let done = false;
</script>

<div class="inline-block select-none bg-emerald-100 rounded">
	<label for={id} class="inline-block px-2 py-1" class:line-through={done}>{text}</label>
	<input type="checkbox" class="mr-2" {id} bind:checked={done} />
</div>
